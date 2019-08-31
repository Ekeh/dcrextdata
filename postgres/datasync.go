package postgres

import (
	"context"
	"fmt"
	"github.com/raedahgroup/dcrextdata/postgres/models"
)

func (pg *PgDb) TableNames() []string {
	return []string{
		models.TableNames.Vote,
		models.TableNames.Block,
		models.TableNames.Mempool,
		models.TableNames.Exchange,
		models.TableNames.ExchangeTick,
		models.TableNames.VSP,
		models.TableNames.VSPTick,
		models.TableNames.PowData,
		models.TableNames.SyncHistory,
	}
}

func (pg *PgDb) LastEntry(ctx context.Context, tableName string, receiver interface{}) error {
	var columnName string
	switch tableName {
	case models.TableNames.Exchange:
		columnName = models.ExchangeColumns.ID
		break
	case models.TableNames.ExchangeTick:
		columnName = models.ExchangeTickColumns.ID
		break
	}

	rows := pg.db.QueryRow(fmt.Sprintf("SELECT MAX(%s) FROM %s", columnName, tableName))
	err := rows.Scan(receiver)
	return err

}
