package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/raedahgroup/dcrextdata/app/helpers"
	"github.com/raedahgroup/dcrextdata/cache"
	"github.com/raedahgroup/dcrextdata/mempool"
	"github.com/raedahgroup/dcrextdata/postgres/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (pg PgDb) MempoolTableName() string {
	return models.TableNames.Mempool
}

func (pg PgDb) BlockTableName() string {
	return models.TableNames.Block
}

func (pg PgDb) VoteTableName() string {
	return models.TableNames.Vote
}

func (pg PgDb) StoreMempool(ctx context.Context, mempoolDto mempool.Mempool) error {
	mempoolModel := mempoolDtoToModel(mempoolDto)
	err := mempoolModel.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if !strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
			return err
		}
	}
	//  tx count 76, total size 54205 B, fees 0.00367100
	log.Infof("Added mempool entry at %s, tx count %2d, total size: %6d B, Total Fee: %010.8f",
		mempoolDto.Time.Format(dateTemplate), mempoolDto.NumberOfTransactions, mempoolDto.Size, mempoolDto.TotalFee)
	return nil
}

func (pg PgDb) StoreMempoolFromSync(ctx context.Context, mempoolDto interface{}) error {
	mempoolModel := mempoolDtoToModel(mempoolDto.(mempool.Mempool))
	err := mempoolModel.Insert(ctx, pg.db, boil.Infer())
	if isUniqueConstraint(err) {
		return nil
	}
	return err
}

func mempoolDtoToModel(mempoolDto mempool.Mempool) models.Mempool {
	return models.Mempool{
		Time:                 mempoolDto.Time,
		FirstSeenTime:        null.TimeFrom(mempoolDto.FirstSeenTime),
		Size:                 null.IntFrom(int(mempoolDto.Size)),
		NumberOfTransactions: null.IntFrom(mempoolDto.NumberOfTransactions),
		Revocations:          null.IntFrom(mempoolDto.Revocations),
		Tickets:              null.IntFrom(mempoolDto.Tickets),
		Voters:               null.IntFrom(mempoolDto.Voters),
		Total:                null.Float64From(mempoolDto.Total),
		TotalFee:             null.Float64From(mempoolDto.TotalFee),
	}
}

func (pg *PgDb) LastMempoolBlockHeight() (height int64, err error) {
	rows := pg.db.QueryRow(lastMempoolBlockHeight)
	err = rows.Scan(&height)
	return
}

func (pg *PgDb) LastMempoolTime() (entryTime time.Time, err error) {
	rows := pg.db.QueryRow(lastMempoolEntryTime)
	err = rows.Scan(&entryTime)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (pg *PgDb) MempoolCount(ctx context.Context) (int64, error) {
	return models.Mempools().Count(ctx, pg.db)
}

func (pg *PgDb) Mempools(ctx context.Context, offtset int, limit int) ([]mempool.Dto, error) {
	mempoolSlice, err := models.Mempools(qm.OrderBy("time DESC"), qm.Offset(offtset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}
	var result []mempool.Dto
	for _, m := range mempoolSlice {
		result = append(result, mempool.Dto{
			TotalFee:             m.TotalFee.Float64,
			FirstSeenTime:        m.FirstSeenTime.Time.Format(dateTemplate),
			Total:                m.Total.Float64,
			Voters:               m.Voters.Int,
			Tickets:              m.Tickets.Int,
			Revocations:          m.Revocations.Int,
			Time:                 m.Time.Format(dateTemplate),
			Size:                 int32(m.Size.Int),
			NumberOfTransactions: m.NumberOfTransactions.Int,
		})
	}
	return result, nil
}

func (pg *PgDb) FetchMempoolForSync(ctx context.Context, date time.Time, offtset int, limit int) ([]mempool.Mempool, int64, error) {
	mempoolSlice, err := models.Mempools(
		models.MempoolWhere.Time.GT(date),
		qm.OrderBy(models.MempoolColumns.Time), qm.Offset(offtset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}
	var result []mempool.Mempool
	for _, m := range mempoolSlice {
		result = append(result, mempool.Mempool{
			TotalFee:             m.TotalFee.Float64,
			FirstSeenTime:        m.FirstSeenTime.Time,
			Total:                m.Total.Float64,
			Voters:               m.Voters.Int,
			Tickets:              m.Tickets.Int,
			Revocations:          m.Revocations.Int,
			Time:                 m.Time,
			Size:                 int32(m.Size.Int),
			NumberOfTransactions: m.NumberOfTransactions.Int,
		})
	}
	totalCount, err := models.Mempools(models.MempoolWhere.Time.GTE(date)).Count(ctx, pg.db)

	return result, totalCount, nil
}

func (pg *PgDb) SaveBlock(ctx context.Context, block mempool.Block) error {
	blockModel := blockDtoToModel(block)
	err := blockModel.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if !strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
			return err
		}
	}

	votes, err := pg.votesByBlock(ctx, int64(block.BlockHeight))
	if err == nil {
		for _, vote := range votes {
			voteModel, err := models.FindVote(ctx, pg.db, vote.Hash)
			if err == nil {
				voteModel.BlockReceiveTime = null.TimeFrom(block.BlockReceiveTime)
				voteModel.BlockHash = null.StringFrom(block.BlockHash)
				_, err = voteModel.Update(ctx, pg.db, boil.Infer())
				if err != nil {
					log.Errorf("Unable to fetch vote for block receive time update: %s", err.Error())
				}
			}
		}
	}

	log.Infof("New block received at %s, PropagationHeight: %d, Hash: ...%s",
		block.BlockReceiveTime.Format(dateMiliTemplate), block.BlockHeight, block.BlockHash[len(block.BlockHash)-23:])
	return nil
}

func (pg *PgDb) SaveBlockFromSync(ctx context.Context, block interface{}) error {
	blockModel := blockDtoToModel(block.(mempool.Block))
	err := blockModel.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if !strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
			return err
		}
	}
	return nil
}

func blockDtoToModel(block mempool.Block) models.Block {
	return models.Block{
		Height:            int(block.BlockHeight),
		Hash:              null.StringFrom(block.BlockHash),
		InternalTimestamp: null.TimeFrom(block.BlockInternalTime),
		ReceiveTime:       null.TimeFrom(block.BlockReceiveTime),
	}
}

func (pg *PgDb) BlockCount(ctx context.Context) (int64, error) {
	return models.Blocks().Count(ctx, pg.db)
}

func (pg *PgDb) Blocks(ctx context.Context, offset int, limit int) ([]mempool.BlockDto, error) {
	blockSlice, err := models.Blocks(qm.OrderBy(fmt.Sprintf("%s DESC", models.BlockColumns.ReceiveTime)),
		qm.Offset(offset), qm.Limit(limit)).All(ctx, pg.db)

	if err != nil {
		return nil, err
	}

	var blocks []mempool.BlockDto

	for _, block := range blockSlice {
		timeDiff := block.ReceiveTime.Time.Sub(block.InternalTimestamp.Time).Seconds()

		votes, err := pg.votesByBlock(ctx, int64(block.Height))
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		blocks = append(blocks, mempool.BlockDto{
			BlockHash:         block.Hash.String,
			BlockHeight:       uint32(block.Height),
			BlockInternalTime: block.InternalTimestamp.Time.Format(dateTemplate),
			BlockReceiveTime:  block.ReceiveTime.Time.Format(dateTemplate),
			Delay:             fmt.Sprintf("%04.2f", timeDiff),
			Votes:             votes,
		})
	}

	return blocks, nil
}

func (pg *PgDb) BlocksWithoutVotes(ctx context.Context, offset int, limit int) ([]mempool.BlockDto, error) {
	blockSlice, err := models.Blocks(qm.OrderBy(fmt.Sprintf("%s DESC", models.BlockColumns.ReceiveTime)), qm.Offset(offset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var blocks []mempool.BlockDto

	for _, block := range blockSlice {
		timeDiff := block.ReceiveTime.Time.Sub(block.InternalTimestamp.Time).Seconds()

		blocks = append(blocks, mempool.BlockDto{
			BlockHash:         block.Hash.String,
			BlockHeight:       uint32(block.Height),
			BlockInternalTime: block.InternalTimestamp.Time.Format(dateTemplate),
			BlockReceiveTime:  block.ReceiveTime.Time.Format(dateTemplate),
			Delay:             fmt.Sprintf("%04.2f", timeDiff),
		})
	}

	return blocks, nil
}

func (pg *PgDb) getBlock(ctx context.Context, height int) (*models.Block, error) {
	block, err := models.Blocks(models.BlockWhere.Height.EQ(height)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (pg *PgDb) FetchBlockForSync(ctx context.Context, blockHeight int64, offtset int, limit int) ([]mempool.Block, int64, error) {
	blockSlice, err := models.Blocks(
		models.BlockWhere.Height.GT(int(blockHeight)),
		qm.OrderBy(models.BlockColumns.ReceiveTime),
		qm.Offset(offtset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}
	var result []mempool.Block
	for _, block := range blockSlice {
		result = append(result, mempool.Block{
			BlockHash:         block.Hash.String,
			BlockHeight:       uint32(block.Height),
			BlockInternalTime: block.InternalTimestamp.Time,
			BlockReceiveTime:  block.ReceiveTime.Time,
		})
	}
	totalCount, err := models.Blocks(models.BlockWhere.Height.GT(int(blockHeight))).Count(ctx, pg.db)

	return result, totalCount, nil
}

func (pg *PgDb) SaveVote(ctx context.Context, vote mempool.Vote) error {
	voteModel := models.Vote{
		Hash:              vote.Hash,
		VotingOn:          null.Int64From(vote.VotingOn),
		BlockHash:         null.StringFrom(vote.BlockHash),
		ReceiveTime:       null.TimeFrom(vote.ReceiveTime),
		TargetedBlockTime: null.TimeFrom(vote.TargetedBlockTime),
		ValidatorID:       null.IntFrom(vote.ValidatorId),
		Validity:          null.StringFrom(vote.Validity),
	}

	// get the target block
	block, err := pg.getBlock(ctx, int(vote.VotingOn))
	if err == nil {
		voteModel.BlockReceiveTime = null.TimeFrom(block.ReceiveTime.Time)
	}

	err = voteModel.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
			return nil
		}
		return err
	}

	log.Infof("New vote received at %s for %d, Validator Id %d, Hash ...%s",
		vote.ReceiveTime.Format(dateMiliTemplate), vote.VotingOn, vote.ValidatorId, vote.Hash[len(vote.Hash)-23:])
	return nil
}

func (pg *PgDb) SaveVoteFromSync(ctx context.Context, voteData interface{}) error {
	vote := voteData.(mempool.Vote)
	voteModel := models.Vote{
		Hash:              vote.Hash,
		VotingOn:          null.Int64From(vote.VotingOn),
		BlockHash:         null.StringFrom(vote.BlockHash),
		ReceiveTime:       null.TimeFrom(vote.ReceiveTime),
		BlockReceiveTime:  null.TimeFrom(vote.BlockReceiveTime),
		TargetedBlockTime: null.TimeFrom(vote.TargetedBlockTime),
		ValidatorID:       null.IntFrom(vote.ValidatorId),
		Validity:          null.StringFrom(vote.Validity),
	}

	err := voteModel.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") { // Ignore duplicate entries
			return nil
		}
	}
	return err
}

func (pg *PgDb) Votes(ctx context.Context, offset int, limit int) ([]mempool.VoteDto, error) {
	voteSlice, err := models.Votes(qm.OrderBy(fmt.Sprintf("%s DESC", models.BlockColumns.ReceiveTime)), qm.Offset(offset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var votes []mempool.VoteDto
	for _, vote := range voteSlice {
		votes = append(votes, pg.voteModelToDto(vote))
	}

	return votes, nil
}

func (pg *PgDb) votesByBlock(ctx context.Context, blockHeight int64) ([]mempool.VoteDto, error) {
	voteSlice, err := models.Votes(models.VoteWhere.VotingOn.EQ(null.Int64From(blockHeight)),
		qm.OrderBy(models.BlockColumns.ReceiveTime)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var votes []mempool.VoteDto
	for _, vote := range voteSlice {
		votes = append(votes, pg.voteModelToDto(vote))
	}

	return votes, nil
}

func (pg *PgDb) voteModelToDto(vote *models.Vote) mempool.VoteDto {
	timeDiff := vote.ReceiveTime.Time.Sub(vote.TargetedBlockTime.Time).Seconds()
	blockReceiveTimeDiff := vote.ReceiveTime.Time.Sub(vote.BlockReceiveTime.Time).Seconds()
	var shortBlockHash string
	if len(vote.BlockHash.String) > 0 {
		shortBlockHash = vote.BlockHash.String[len(vote.BlockHash.String)-8:]
	}

	return mempool.VoteDto{
		Hash:                  vote.Hash,
		ReceiveTime:           vote.ReceiveTime.Time.Format(dateTemplate),
		TargetedBlockTimeDiff: fmt.Sprintf("%04.2f", timeDiff),
		BlockReceiveTimeDiff:  fmt.Sprintf("%04.2f", blockReceiveTimeDiff),
		VotingOn:              vote.VotingOn.Int64,
		BlockHash:             vote.BlockHash.String,
		ShortBlockHash:        shortBlockHash,
		ValidatorId:           vote.ValidatorID.Int,
		Validity:              vote.Validity.String,
	}
}

func (pg *PgDb) VotesCount(ctx context.Context) (int64, error) {
	return models.Votes().Count(ctx, pg.db)
}

func (pg *PgDb) FetchVoteForSync(ctx context.Context, date time.Time, offtset int, limit int) ([]mempool.Vote, int64, error) {
	voteSlices, err := models.Votes(
		models.VoteWhere.ReceiveTime.GTE(null.TimeFrom(date)),
		qm.OrderBy(models.VoteColumns.ReceiveTime),
		qm.Offset(offtset), qm.Limit(limit)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}
	var result []mempool.Vote
	for _, vote := range voteSlices {
		result = append(result, mempool.Vote{
			Hash:              vote.Hash,
			ReceiveTime:       vote.ReceiveTime.Time,
			TargetedBlockTime: vote.TargetedBlockTime.Time,
			BlockReceiveTime:  vote.BlockReceiveTime.Time,
			VotingOn:          vote.VotingOn.Int64,
			BlockHash:         vote.BlockHash.String,
			ValidatorId:       vote.ValidatorID.Int,
			Validity:          vote.Validity.String,
		})
	}
	totalCount, err := models.Votes(models.VoteWhere.ReceiveTime.GTE(null.TimeFrom(date))).Count(ctx, pg.db)

	return result, totalCount, nil
}

func (pg *PgDb) propagationVoteChartDataByHeight(ctx context.Context, height int32) ([]mempool.PropagationChartData, error) {
	voteSlice, err := models.Votes(
		models.VoteWhere.VotingOn.GT(null.Int64From(int64(height))),
		qm.OrderBy(models.VoteColumns.VotingOn)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var chartData []mempool.PropagationChartData
	for _, vote := range voteSlice {
		blockReceiveTimeDiff := vote.ReceiveTime.Time.Sub(vote.BlockReceiveTime.Time).Seconds()
		chartData = append(chartData, mempool.PropagationChartData{
			BlockHeight: vote.VotingOn.Int64, TimeDifference: blockReceiveTimeDiff,
		})
	}

	return chartData, nil
}

func (pg *PgDb) propagationBlockChartData(ctx context.Context, height int) ([]mempool.PropagationChartData, error) {
	blockSlice, err := models.Blocks(
		models.BlockWhere.Height.GT(height),
		qm.OrderBy(models.BlockColumns.Height)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var chartData []mempool.PropagationChartData
	for _, block := range blockSlice {
		blockReceiveTimeDiff := block.ReceiveTime.Time.Sub(block.InternalTimestamp.Time).Seconds()
		chartData = append(chartData, mempool.PropagationChartData{
			BlockHeight: int64(block.Height), TimeDifference: blockReceiveTimeDiff,
		})
	}

	return chartData, nil
}

func (pg *PgDb) fetchBlockReceiveTimeByHeight(ctx context.Context, height int32) ([]mempool.BlockReceiveTime, error) {
	blockSlice, err := models.Blocks(
		models.BlockWhere.Height.GT(int(height)),
		qm.Select(models.BlockColumns.Height, models.BlockColumns.ReceiveTime),
		qm.OrderBy(models.BlockColumns.Height),
	).All(ctx, pg.db)

	if err != nil {
		return nil, err
	}

	var chartData []mempool.BlockReceiveTime
	for _, block := range blockSlice {
		chartData = append(chartData, mempool.BlockReceiveTime{
			BlockHeight: int64(block.Height),
			ReceiveTime: block.ReceiveTime.Time,
		})
	}

	return chartData, nil
}

// *****CHARTS******* //

func (pg *PgDb) retrieveChartMempool(ctx context.Context, charts *cache.ChartData) (interface{}, func(), error) {
	ctx, cancel := context.WithTimeout(ctx, pg.queryTimeout)

	charts.PropagationHeight()
	mempoolSlice, err := models.Mempools(models.MempoolWhere.Time.GT(helpers.UnixTime(int64(charts.MempoolTime())))).All(ctx, pg.db)
	if err != nil {
		return nil, cancel, fmt.Errorf("chartBlocks: %s", err.Error())
	}
	return mempoolSlice, cancel, nil
}

// Append the results from retrieveChartMempool to the provided ChartData.
// This is the Appender half of a pair that make up a cache.ChartUpdater.
func appendChartMempool(charts *cache.ChartData, mempoolSliceInt interface{}) error {
	mempoolSlice := mempoolSliceInt.(models.MempoolSlice)
	chartsMempool := charts.Mempool
	for _, mempoolData := range mempoolSlice {
		chartsMempool.Time = append(chartsMempool.Time, uint64(mempoolData.Time.UTC().Unix()))
		chartsMempool.Fees = append(chartsMempool.Fees, mempoolData.TotalFee.Float64)
		chartsMempool.TxCount = append(chartsMempool.TxCount, uint64(mempoolData.NumberOfTransactions.Int))
		chartsMempool.Size = append(chartsMempool.Size, uint64(mempoolData.Size.Int))
	}
	return nil
}

type propagationSet struct {
	height                    []uint64
	blockDelay                []float64
	voteReceiveTimeDeviations []float64
	blockPropagation          map[string][]float64
}

func (pg *PgDb) fetchBlockPropagationChart(ctx context.Context, charts *cache.ChartData) (interface{}, func(), error) {
	emptyCancelFunc := func() {}
	var propagationSet propagationSet

	chartsBlockHeight := charts.PropagationHeight()
	blockDelays, err := pg.propagationBlockChartData(ctx, int(chartsBlockHeight))
	if err != nil && err != sql.ErrNoRows {
		return nil, emptyCancelFunc, err
	}

	localBlockReceiveTime := make(map[uint64]float64)
	for _, record := range blockDelays {
		propagationSet.height = append(propagationSet.height, uint64(record.BlockHeight))
		timeDifference, _ := strconv.ParseFloat(fmt.Sprintf("%04.2f", record.TimeDifference), 64)
		propagationSet.blockDelay = append(propagationSet.blockDelay, timeDifference)

		localBlockReceiveTime[uint64(record.BlockHeight)] = timeDifference
	}

	votesReceiveTime, err := pg.propagationVoteChartDataByHeight(ctx, chartsBlockHeight)
	if err != nil && err != sql.ErrNoRows {
		return nil, emptyCancelFunc, err
	}
	var votesTimeDeviations = make(map[int64][]float64)

	for _, record := range votesReceiveTime {
		votesTimeDeviations[record.BlockHeight] = append(votesTimeDeviations[record.BlockHeight], record.TimeDifference)
	}

	for _, height := range propagationSet.height {
		if deviations, found := votesTimeDeviations[int64(height)]; found {
			var totalTime float64
			for _, timeDiff := range deviations {
				totalTime += timeDiff
			}
			timeDifference, _ := strconv.ParseFloat(fmt.Sprintf("%04.2f", totalTime/float64(len(deviations))*1000), 64)
			propagationSet.voteReceiveTimeDeviations = append(propagationSet.voteReceiveTimeDeviations, timeDifference)
			continue
		}
		propagationSet.voteReceiveTimeDeviations = append(propagationSet.voteReceiveTimeDeviations, 0)
	}

	propagationSet.blockPropagation = make(map[string][]float64)
	for _, source := range pg.syncSources {
		db, err := pg.syncSourceDbProvider(source)
		if err != nil {
			return nil, emptyCancelFunc, err
		}

		blockDelays, err := db.propagationBlockChartData(ctx, int(chartsBlockHeight))
		if err != nil && err != sql.ErrNoRows {
			return nil, emptyCancelFunc, err
		}

		receiveTimeMap := make(map[uint64]float64)
		for _, record := range blockDelays {

			receiveTimeMap[uint64(record.BlockHeight)], _ = strconv.ParseFloat(fmt.Sprintf("%04.2f", record.TimeDifference), 64)
		}

		for _, height := range propagationSet.height {
			if sourceTime, found := receiveTimeMap[height]; found {
				propagationSet.blockPropagation[source] = append(propagationSet.blockPropagation[source], localBlockReceiveTime[height]-sourceTime)
				continue
			}
			propagationSet.blockPropagation[source] = append(propagationSet.blockPropagation[source], 0)
		}
	}

	return propagationSet, emptyCancelFunc, nil
}

func appendBlockPropagationChart(charts *cache.ChartData, data interface{}) error {
	propagationSet := data.(propagationSet)
	for _, height := range propagationSet.height {
		charts.Propagation.Height = append(charts.Propagation.Height, height)
	}
	for _, delay := range propagationSet.blockDelay {
		charts.Propagation.BlockDelays = append(charts.Propagation.BlockDelays, delay)
	}
	for _, voteTime := range propagationSet.voteReceiveTimeDeviations {
		charts.Propagation.VotesReceiveTimeDeviations = append(charts.Propagation.VotesReceiveTimeDeviations, voteTime)
	}

	for source, deviations := range propagationSet.blockPropagation {
		charts.Propagation.BlockPropagation[source] = append(charts.Propagation.BlockPropagation[source], deviations...)
	}

	return nil
}
