// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package mempool

import (
	"context"
	"sync"
	"time"

	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/chaincfg/chainhash"
	"github.com/decred/dcrd/dcrjson"
	"github.com/decred/dcrd/rpcclient"
	"github.com/decred/dcrd/wire"
	exptypes "github.com/decred/dcrdata/explorer/types"
	"github.com/decred/dcrdata/txhelpers/v2"
)

func NewCollector(config *rpcclient.ConnConfig, activeChain *chaincfg.Params, dataStore DataStore) *Collector {
	return &Collector{
		dcrdClientConfig: config,
		dataStore:        dataStore,
		activeChain:      activeChain,
	}
}

func (c *Collector) StartMonitoring(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var ticketIndsMutex sync.Mutex
	ticketInds := make(exptypes.BlockValidatorIndex)

	freeClient, err := rpcclient.New(c.dcrdClientConfig, nil)
	if err != nil {
		log.Errorf("Error in opening a dcrd connection: %s", err.Error())
		return
	}
	defer freeClient.Shutdown()

	ntfnHandlers := rpcclient.NotificationHandlers{
		OnTxAcceptedVerbose: func(txDetails *dcrjson.TxRawResult) {
			go func() {
				if !c.syncIsDone {
					return
				}
				receiveTime := time.Now()

				msgTx, err := txhelpers.MsgTxFromHex(txDetails.Hex)
				if err != nil {
					log.Errorf("Failed to decode transaction hex: %v", err)
					return
				}

				if txType := txhelpers.DetermineTxTypeString(msgTx); txType != "Vote" {
					return
				}

				var voteInfo *exptypes.VoteInfo
				validation, version, bits, choices, err := txhelpers.SSGenVoteChoices(msgTx, c.activeChain)
				if err != nil {
					log.Errorf("Error in getting vote choice: %s", err.Error())
					return
				}

				voteInfo = &exptypes.VoteInfo{
					Validation: exptypes.BlockValidation{
						Hash:     validation.Hash.String(),
						Height:   validation.Height,
						Validity: validation.Validity,
					},
					Version:     version,
					Bits:        bits,
					Choices:     choices,
					TicketSpent: msgTx.TxIn[1].PreviousOutPoint.Hash.String(),
				}

				ticketIndsMutex.Lock()
				voteInfo.SetTicketIndex(ticketInds)
				ticketIndsMutex.Unlock()

				vote := Vote{
					ReceiveTime: receiveTime,
					VotingOn:    validation.Height,
					Hash:        txDetails.Txid,
					ValidatorId: voteInfo.MempoolTicketIndex,
				}

				// wait for some time for the block to get added to the blockchain
				time.Sleep(2 * time.Second)

				targetedBlock, err := freeClient.GetBlock(&validation.Hash)
				if err != nil {
					log.Errorf("Error in getting validation targeted block: %s", err.Error())
					return
				}

				vote.TargetedBlockTime = targetedBlock.Header.Timestamp

				if err = c.dataStore.SaveVote(ctx, vote); err != nil {
					log.Error(err)
				}
			}()
		},

		OnBlockConnected: func(blockHeaderSerialized []byte, transactions [][]byte) {
			if !c.syncIsDone {
				return
			}
			blockHeader := new(wire.BlockHeader)
			err := blockHeader.FromBytes(blockHeaderSerialized)
			if err != nil {
				log.Error("Failed to deserialize blockHeader in new block notification: %v", err)
				return
			}

			block := Block{
				BlockInternalTime: blockHeader.Timestamp,
				BlockReceiveTime:  time.Now(),
				BlockHash:         blockHeader.BlockHash().String(),
				BlockHeight:       blockHeader.Height,
			}
			if err = c.dataStore.SaveBlock(ctx, block); err != nil {
				log.Error(err)
			}
		},
	}

	client, err := rpcclient.New(c.dcrdClientConfig, &ntfnHandlers)
	if err != nil {
		log.Errorf("Error in opening a dcrd connection: %s", err.Error())
		return
	}

	defer client.Shutdown()

	if err := client.NotifyNewTransactions(true); err != nil {
		log.Error(err)
	}

	if err := client.NotifyBlocks(); err != nil {
		log.Error(err)
	}

	var mu sync.Mutex

	collectMempool := func() {
		mu.Lock()
		defer mu.Unlock()

		mempoolTransactionMap, err := client.GetRawMempoolVerbose(dcrjson.GRMAll)
		if err != nil {
			log.Error(err)
			return
		}

		if len(mempoolTransactionMap) == 0 {
			return
		}

		// there wont be transactions in the mempool while sync is going on
		c.syncIsDone = true // todo: we need a better way to determine the sync status of dcrd

		mempoolDto := Mempool{
			NumberOfTransactions: len(mempoolTransactionMap),
			Time:                 time.Now(),
			FirstSeenTime:        time.Now(), //todo: use the time of the first tx in the mempool
		}

		for hashString, tx := range mempoolTransactionMap {
			hash, err := chainhash.NewHashFromStr(hashString)
			if err != nil {
				log.Error(err)
				continue
			}
			rawTx, err := client.GetRawTransactionVerbose(hash)
			if err != nil {
				log.Error(err)
				continue
			}

			totalOut := 0.0
			for _, v := range rawTx.Vout {
				totalOut += v.Value
			}

			mempoolDto.Total += totalOut
			mempoolDto.TotalFee += tx.Fee
			mempoolDto.Size += tx.Size
			if mempoolDto.FirstSeenTime.Unix() > tx.Time {
				mempoolDto.FirstSeenTime = time.Unix(tx.Time, 0)
			}

		}

		votes, err := client.GetRawMempool(dcrjson.GRMVotes)
		if err != nil {
			log.Error(err)
			return
		}
		mempoolDto.Voters = len(votes)

		tickets, err := client.GetRawMempool(dcrjson.GRMTickets)
		if err != nil {
			log.Error(err)
			return
		}
		mempoolDto.Tickets = len(tickets)

		revocations, err := client.GetRawMempool(dcrjson.GRMRevocations)
		if err != nil {
			log.Error(err)
			return
		}
		mempoolDto.Revocations = len(revocations)

		err = c.dataStore.StoreMempool(ctx, mempoolDto)
		if err != nil {
			log.Error(err)
		}
	}

	collectMempool()
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			collectMempool()
			break
		case <-ctx.Done():
			return
		}
	}
}