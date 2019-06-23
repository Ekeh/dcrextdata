package postgres

import "fmt"

const (
	createExchangeTable = `CREATE TABLE IF NOT EXISTS exchange (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL);`

	createExchangeTickTable = `CREATE TABLE IF NOT EXISTS exchange_tick (
		id SERIAL PRIMARY KEY,
		exchange_id INT REFERENCES exchange(id) NOT NULL, 
		interval INT NOT NULL,
		high FLOAT NOT NULL,
		low FLOAT NOT NULL,
		open FLOAT NOT NULL,
		close FLOAT NOT NULL,
		volume FLOAT NOT NULL,
		currency_pair TEXT NOT NULL,
		time TIMESTAMPTZ NOT NULL
	);`

	createExchangeTickIndex = `CREATE UNIQUE INDEX IF NOT EXISTS exchange_tick_idx ON exchange_tick (exchange_id, interval, currency_pair, time);`

	createVSPInfoTable = `CREATE TABLE IF NOT EXISTS vsp (
		id SERIAL PRIMARY KEY,
		name TEXT,
		api_enabled BOOLEAN,
		api_versions_supported INT8[],
		network TEXT,
		url TEXT,
		launched TIMESTAMPTZ
	);`

	createVSPTickTable = `CREATE TABLE IF NOT EXISTS vsp_tick (
		id SERIAL PRIMARY KEY,
		vsp_id INT REFERENCES vsp(id) NOT NULL,
		immature INT NOT NULL,
		live INT NOT NULL,
		voted INT NOT NULL,
		missed INT NOT NULL,
		pool_fees FLOAT NOT NULL,
		proportion_live FLOAT NOT NULL,
		proportion_missed FLOAT NOT NULL,
		user_count INT NOT NULL,
		users_active INT NOT NULL,
		time TIMESTAMPTZ NOT NULL
	);`

	createVSPTickIndex = `CREATE UNIQUE INDEX IF NOT EXISTS vsp_tick_idx ON vsp_tick (vsp_id,immature,live,voted,missed,pool_fees,proportion_live,proportion_missed,user_count,users_active, time);`

	// PoW table
	createPowDataTable = `CREATE TABLE IF NOT EXISTS pow_data (
		time INT,
		network_hashrate VARCHAR(25),
		pool_hashrate VARCHAR(25),
		workers INT,
		network_difficulty FLOAT8,
		coin_price VARCHAR(25),
		btc_price VARCHAR(25),
		source VARCHAR(25),
		PRIMARY KEY (time, source)
	);`

	lastPowEntryTime = `SELECT time FROM pow_data WHERE source=$1 ORDER BY time DESC LIMIT 1`

	createMempoolTable = `CREATE TABLE IF NOT EXISTS mempool (
		first_seen_time INT8,
		number_of_transactions INT,
		size INT,
		block_receive_time INT8,
		block_internal_time INT8,
		block_height INT,
		block_hash VARCHAR(128),
		PRIMARY KEY (block_height,block_hash)
	);`

	lastMempoolBlockHeight = `SELECT last_block_height FROM mempool ORDER BY last_block_height DESC LIMIT 1`
)

func (pg *PgDb) CreateExchangeTable() error {
	log.Trace("Creating exchange tick table")
	_, err := pg.db.Exec(createExchangeTable)
	return err
}

func (pg *PgDb) ExchangeTableExits() bool {
	exists, _ := pg.tableExists("exchange")
	return exists
}

func (pg *PgDb) CreateExchangeTickTable() error {
	log.Trace("Creating exchange tick table")
	_, err := pg.db.Exec(createExchangeTickTable)
	return err
}

func (pg *PgDb) CreateExchangeTickIndex() error {
	_, err := pg.db.Exec(createExchangeTickIndex)
	return err
}

func (pg *PgDb) ExchangeTickTableExits() bool {
	exists, _ := pg.tableExists("exchange_tick")
	return exists
}

func (pg *PgDb) CreateVSPInfoTables() error {
	_, err := pg.db.Exec(createVSPInfoTable)
	return err
}

func (pg *PgDb) VSPInfoTableExits() bool {
	exists, _ := pg.tableExists("vsp")
	return exists
}

func (pg *PgDb) CreateVSPTickTables() error {
	_, err := pg.db.Exec(createVSPTickTable)
	return err
}

func (pg *PgDb) CreateVSPTickIndex() error {
	_, err := pg.db.Exec(createVSPTickIndex)

	return err
}

func (pg *PgDb) VSPTickTableExits() bool {
	exists, _ := pg.tableExists("vsp_tick")
	return exists
}

func (pg *PgDb) CreatePowDataTable() error {
	_, err := pg.db.Exec(createPowDataTable)
	return err
}

func (pg *PgDb) PowDataTableExits() bool {
	exists, _ := pg.tableExists("pow_data")
	return exists
}

func (pg *PgDb) CreateMempoolDataTable() error {
	_, err := pg.db.Exec(createMempoolTable)
	return err
}

func (pg *PgDb) MempoolDataTableExits() bool {
	exists, _ := pg.tableExists("mempool")
	return exists
}

func (pg *PgDb) tableExists(name string) (bool, error) {
	rows, err := pg.db.Query(`SELECT relname FROM pg_class WHERE relname = $1`, name)
	if err == nil {
		defer func() {
			if e := rows.Close(); e != nil {
				log.Error("Close of Query failed: ", e)
			}
		}()
		return rows.Next(), nil
	}
	return false, err
}

func (pg *PgDb) DropAllTables() error {
	// vsp_tick
	if err := pg.dropIndex("vsp_tick_idx"); err != nil {
		return err
	}

	if err := pg.dropTable("vsp_tick"); err != nil {
		return err
	}

	// vsp
	if err := pg.dropTable("vsp"); err != nil {
		return err
	}

	// exchange_tick
	if err := pg.dropIndex("exchange_tick_idx"); err != nil {
		return err
	}

	if err := pg.dropTable("exchange_tick"); err != nil {
		return err
	}

	// exchange
	if err := pg.dropTable("exchange"); err != nil {
		return err
	}

	// pow_data
	if err := pg.dropTable("pow_data"); err != nil {
		return err
	}

	// mempool
	if err := pg.dropTable("mempool"); err != nil {
		return err
	}

	// pow_data
	return nil
}

func (pg *PgDb) dropTable(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, name))
	return err
}

func (pg *PgDb) dropIndex(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP INDEX IF EXISTS %s;`, name))
	return err
}
