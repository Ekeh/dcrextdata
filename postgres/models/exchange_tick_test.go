// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testExchangeTicks(t *testing.T) {
	t.Parallel()

	query := ExchangeTicks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testExchangeTicksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangeTicksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ExchangeTicks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangeTicksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExchangeTickSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testExchangeTicksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ExchangeTickExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ExchangeTick exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ExchangeTickExists to return true, but got false.")
	}
}

func testExchangeTicksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	exchangeTickFound, err := FindExchangeTick(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if exchangeTickFound == nil {
		t.Error("want a record, got nil")
	}
}

func testExchangeTicksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ExchangeTicks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testExchangeTicksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ExchangeTicks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testExchangeTicksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	exchangeTickOne := &ExchangeTick{}
	exchangeTickTwo := &ExchangeTick{}
	if err = randomize.Struct(seed, exchangeTickOne, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}
	if err = randomize.Struct(seed, exchangeTickTwo, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = exchangeTickOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = exchangeTickTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ExchangeTicks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testExchangeTicksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	exchangeTickOne := &ExchangeTick{}
	exchangeTickTwo := &ExchangeTick{}
	if err = randomize.Struct(seed, exchangeTickOne, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}
	if err = randomize.Struct(seed, exchangeTickTwo, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = exchangeTickOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = exchangeTickTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testExchangeTicksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExchangeTicksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(exchangeTickColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testExchangeTickToOneExchangeUsingExchange(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ExchangeTick
	var foreign Exchange

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, exchangeTickDBTypes, false, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ExchangeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Exchange().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ExchangeTickSlice{&local}
	if err = local.L.LoadExchange(ctx, tx, false, (*[]*ExchangeTick)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Exchange = nil
	if err = local.L.LoadExchange(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testExchangeTickToOneSetOpExchangeUsingExchange(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ExchangeTick
	var b, c Exchange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, exchangeTickDBTypes, false, strmangle.SetComplement(exchangeTickPrimaryKeyColumns, exchangeTickColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Exchange{&b, &c} {
		err = a.SetExchange(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Exchange != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ExchangeTicks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ExchangeID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ExchangeID))
		reflect.Indirect(reflect.ValueOf(&a.ExchangeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ExchangeID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeID, x.ID)
		}
	}
}

func testExchangeTicksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExchangeTicksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ExchangeTickSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testExchangeTicksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ExchangeTicks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	exchangeTickDBTypes = map[string]string{`ID`: `integer`, `ExchangeID`: `integer`, `Interval`: `integer`, `High`: `double precision`, `Low`: `double precision`, `Open`: `double precision`, `Close`: `double precision`, `Volume`: `double precision`, `CurrencyPair`: `text`, `Time`: `timestamp with time zone`}
	_                   = bytes.MinRead
)

func testExchangeTicksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(exchangeTickPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(exchangeTickAllColumns) == len(exchangeTickPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testExchangeTicksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(exchangeTickAllColumns) == len(exchangeTickPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ExchangeTick{}
	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, exchangeTickDBTypes, true, exchangeTickPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(exchangeTickAllColumns, exchangeTickPrimaryKeyColumns) {
		fields = exchangeTickAllColumns
	} else {
		fields = strmangle.SetComplement(
			exchangeTickAllColumns,
			exchangeTickPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ExchangeTickSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testExchangeTicksUpsert(t *testing.T) {
	t.Parallel()

	if len(exchangeTickAllColumns) == len(exchangeTickPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ExchangeTick{}
	if err = randomize.Struct(seed, &o, exchangeTickDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ExchangeTick: %s", err)
	}

	count, err := ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, exchangeTickDBTypes, false, exchangeTickPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ExchangeTick struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ExchangeTick: %s", err)
	}

	count, err = ExchangeTicks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
