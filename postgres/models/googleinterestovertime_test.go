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

func testGoogleinterestovertimes(t *testing.T) {
	t.Parallel()

	query := Googleinterestovertimes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGoogleinterestovertimesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
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

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoogleinterestovertimesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Googleinterestovertimes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoogleinterestovertimesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoogleinterestovertimeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoogleinterestovertimesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GoogleinterestovertimeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Googleinterestovertime exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GoogleinterestovertimeExists to return true, but got false.")
	}
}

func testGoogleinterestovertimesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	googleinterestovertimeFound, err := FindGoogleinterestovertime(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if googleinterestovertimeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGoogleinterestovertimesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Googleinterestovertimes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGoogleinterestovertimesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Googleinterestovertimes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGoogleinterestovertimesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	googleinterestovertimeOne := &Googleinterestovertime{}
	googleinterestovertimeTwo := &Googleinterestovertime{}
	if err = randomize.Struct(seed, googleinterestovertimeOne, googleinterestovertimeDBTypes, false, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}
	if err = randomize.Struct(seed, googleinterestovertimeTwo, googleinterestovertimeDBTypes, false, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = googleinterestovertimeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = googleinterestovertimeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Googleinterestovertimes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGoogleinterestovertimesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	googleinterestovertimeOne := &Googleinterestovertime{}
	googleinterestovertimeTwo := &Googleinterestovertime{}
	if err = randomize.Struct(seed, googleinterestovertimeOne, googleinterestovertimeDBTypes, false, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}
	if err = randomize.Struct(seed, googleinterestovertimeTwo, googleinterestovertimeDBTypes, false, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = googleinterestovertimeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = googleinterestovertimeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testGoogleinterestovertimesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoogleinterestovertimesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(googleinterestovertimeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoogleinterestovertimesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
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

func testGoogleinterestovertimesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoogleinterestovertimeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGoogleinterestovertimesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Googleinterestovertimes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	googleinterestovertimeDBTypes = map[string]string{`ID`: `integer`, `Geo`: `character varying`, `FormattedTime`: `character varying`, `FormattedAxisTime`: `character varying`, `Value`: `integer`, `Keyword`: `character varying`, `Date`: `timestamp without time zone`}
	_                             = bytes.MinRead
)

func testGoogleinterestovertimesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(googleinterestovertimePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(googleinterestovertimeAllColumns) == len(googleinterestovertimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGoogleinterestovertimesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(googleinterestovertimeAllColumns) == len(googleinterestovertimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Googleinterestovertime{}
	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, googleinterestovertimeDBTypes, true, googleinterestovertimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(googleinterestovertimeAllColumns, googleinterestovertimePrimaryKeyColumns) {
		fields = googleinterestovertimeAllColumns
	} else {
		fields = strmangle.SetComplement(
			googleinterestovertimeAllColumns,
			googleinterestovertimePrimaryKeyColumns,
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

	slice := GoogleinterestovertimeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGoogleinterestovertimesUpsert(t *testing.T) {
	t.Parallel()

	if len(googleinterestovertimeAllColumns) == len(googleinterestovertimePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Googleinterestovertime{}
	if err = randomize.Struct(seed, &o, googleinterestovertimeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Googleinterestovertime: %s", err)
	}

	count, err := Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, googleinterestovertimeDBTypes, false, googleinterestovertimePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Googleinterestovertime struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Googleinterestovertime: %s", err)
	}

	count, err = Googleinterestovertimes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
