// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PowDatum is an object representing the database table.
type PowDatum struct {
	Time              int          `boil:"time" json:"time" toml:"time" yaml:"time"`
	NetworkHashrate   null.String  `boil:"network_hashrate" json:"network_hashrate,omitempty" toml:"network_hashrate" yaml:"network_hashrate,omitempty"`
	PoolHashrate      null.String  `boil:"pool_hashrate" json:"pool_hashrate,omitempty" toml:"pool_hashrate" yaml:"pool_hashrate,omitempty"`
	Workers           null.Int     `boil:"workers" json:"workers,omitempty" toml:"workers" yaml:"workers,omitempty"`
	NetworkDifficulty null.Float64 `boil:"network_difficulty" json:"network_difficulty,omitempty" toml:"network_difficulty" yaml:"network_difficulty,omitempty"`
	CoinPrice         null.String  `boil:"coin_price" json:"coin_price,omitempty" toml:"coin_price" yaml:"coin_price,omitempty"`
	BTCPrice          null.String  `boil:"btc_price" json:"btc_price,omitempty" toml:"btc_price" yaml:"btc_price,omitempty"`
	Source            string       `boil:"source" json:"source" toml:"source" yaml:"source"`

	R *powDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L powDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PowDatumColumns = struct {
	Time              string
	NetworkHashrate   string
	PoolHashrate      string
	Workers           string
	NetworkDifficulty string
	CoinPrice         string
	BTCPrice          string
	Source            string
}{
	Time:              "time",
	NetworkHashrate:   "network_hashrate",
	PoolHashrate:      "pool_hashrate",
	Workers:           "workers",
	NetworkDifficulty: "network_difficulty",
	CoinPrice:         "coin_price",
	BTCPrice:          "btc_price",
	Source:            "source",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PowDatumWhere = struct {
	Time              whereHelperint
	NetworkHashrate   whereHelpernull_String
	PoolHashrate      whereHelpernull_String
	Workers           whereHelpernull_Int
	NetworkDifficulty whereHelpernull_Float64
	CoinPrice         whereHelpernull_String
	BTCPrice          whereHelpernull_String
	Source            whereHelperstring
}{
	Time:              whereHelperint{field: "\"pow_data\".\"time\""},
	NetworkHashrate:   whereHelpernull_String{field: "\"pow_data\".\"network_hashrate\""},
	PoolHashrate:      whereHelpernull_String{field: "\"pow_data\".\"pool_hashrate\""},
	Workers:           whereHelpernull_Int{field: "\"pow_data\".\"workers\""},
	NetworkDifficulty: whereHelpernull_Float64{field: "\"pow_data\".\"network_difficulty\""},
	CoinPrice:         whereHelpernull_String{field: "\"pow_data\".\"coin_price\""},
	BTCPrice:          whereHelpernull_String{field: "\"pow_data\".\"btc_price\""},
	Source:            whereHelperstring{field: "\"pow_data\".\"source\""},
}

// PowDatumRels is where relationship names are stored.
var PowDatumRels = struct {
}{}

// powDatumR is where relationships are stored.
type powDatumR struct {
}

// NewStruct creates a new relationship struct
func (*powDatumR) NewStruct() *powDatumR {
	return &powDatumR{}
}

// powDatumL is where Load methods for each relationship are stored.
type powDatumL struct{}

var (
	powDatumAllColumns            = []string{"time", "network_hashrate", "pool_hashrate", "workers", "network_difficulty", "coin_price", "btc_price", "source"}
	powDatumColumnsWithoutDefault = []string{"time", "network_hashrate", "pool_hashrate", "workers", "network_difficulty", "coin_price", "btc_price", "source"}
	powDatumColumnsWithDefault    = []string{}
	powDatumPrimaryKeyColumns     = []string{"time", "source"}
)

type (
	// PowDatumSlice is an alias for a slice of pointers to PowDatum.
	// This should generally be used opposed to []PowDatum.
	PowDatumSlice []*PowDatum

	powDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	powDatumType                 = reflect.TypeOf(&PowDatum{})
	powDatumMapping              = queries.MakeStructMapping(powDatumType)
	powDatumPrimaryKeyMapping, _ = queries.BindMapping(powDatumType, powDatumMapping, powDatumPrimaryKeyColumns)
	powDatumInsertCacheMut       sync.RWMutex
	powDatumInsertCache          = make(map[string]insertCache)
	powDatumUpdateCacheMut       sync.RWMutex
	powDatumUpdateCache          = make(map[string]updateCache)
	powDatumUpsertCacheMut       sync.RWMutex
	powDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single powDatum record from the query.
func (q powDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PowDatum, error) {
	o := &PowDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pow_data")
	}

	return o, nil
}

// All returns all PowDatum records from the query.
func (q powDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (PowDatumSlice, error) {
	var o []*PowDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PowDatum slice")
	}

	return o, nil
}

// Count returns the count of all PowDatum records in the query.
func (q powDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pow_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q powDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pow_data exists")
	}

	return count > 0, nil
}

// PowData retrieves all the records using an executor.
func PowData(mods ...qm.QueryMod) powDatumQuery {
	mods = append(mods, qm.From("\"pow_data\""))
	return powDatumQuery{NewQuery(mods...)}
}

// FindPowDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPowDatum(ctx context.Context, exec boil.ContextExecutor, time int, source string, selectCols ...string) (*PowDatum, error) {
	powDatumObj := &PowDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pow_data\" where \"time\"=$1 AND \"source\"=$2", sel,
	)

	q := queries.Raw(query, time, source)

	err := q.Bind(ctx, exec, powDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pow_data")
	}

	return powDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PowDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pow_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(powDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	powDatumInsertCacheMut.RLock()
	cache, cached := powDatumInsertCache[key]
	powDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			powDatumAllColumns,
			powDatumColumnsWithDefault,
			powDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(powDatumType, powDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(powDatumType, powDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pow_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pow_data\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pow_data")
	}

	if !cached {
		powDatumInsertCacheMut.Lock()
		powDatumInsertCache[key] = cache
		powDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PowDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PowDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	powDatumUpdateCacheMut.RLock()
	cache, cached := powDatumUpdateCache[key]
	powDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			powDatumAllColumns,
			powDatumPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pow_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pow_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, powDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(powDatumType, powDatumMapping, append(wl, powDatumPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update pow_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pow_data")
	}

	if !cached {
		powDatumUpdateCacheMut.Lock()
		powDatumUpdateCache[key] = cache
		powDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q powDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pow_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pow_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PowDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pow_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, powDatumPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in powDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all powDatum")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PowDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pow_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(powDatumColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	powDatumUpsertCacheMut.RLock()
	cache, cached := powDatumUpsertCache[key]
	powDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			powDatumAllColumns,
			powDatumColumnsWithDefault,
			powDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			powDatumAllColumns,
			powDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pow_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(powDatumPrimaryKeyColumns))
			copy(conflict, powDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pow_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(powDatumType, powDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(powDatumType, powDatumMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pow_data")
	}

	if !cached {
		powDatumUpsertCacheMut.Lock()
		powDatumUpsertCache[key] = cache
		powDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PowDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PowDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PowDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), powDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"pow_data\" WHERE \"time\"=$1 AND \"source\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pow_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pow_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q powDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no powDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pow_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pow_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PowDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pow_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powDatumPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from powDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pow_data")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PowDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPowDatum(ctx, exec, o.Time, o.Source)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PowDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PowDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pow_data\".* FROM \"pow_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PowDatumSlice")
	}

	*o = slice

	return nil
}

// PowDatumExists checks if the PowDatum row exists.
func PowDatumExists(ctx context.Context, exec boil.ContextExecutor, time int, source string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pow_data\" where \"time\"=$1 AND \"source\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, time, source)
	}

	row := exec.QueryRowContext(ctx, sql, time, source)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pow_data exists")
	}

	return exists, nil
}
