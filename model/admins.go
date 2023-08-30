// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Admin is an object representing the database table.
type Admin struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name     string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SteamID  int64  `boil:"steam_id" json:"steam_id" toml:"steam_id" yaml:"steam_id"`
	Immunity int    `boil:"immunity" json:"immunity" toml:"immunity" yaml:"immunity"`
	Flags    string `boil:"flags" json:"flags" toml:"flags" yaml:"flags"`

	R *adminR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminColumns = struct {
	ID       string
	Name     string
	SteamID  string
	Immunity string
	Flags    string
}{
	ID:       "id",
	Name:     "name",
	SteamID:  "steam_id",
	Immunity: "immunity",
	Flags:    "flags",
}

var AdminTableColumns = struct {
	ID       string
	Name     string
	SteamID  string
	Immunity string
	Flags    string
}{
	ID:       "admins.id",
	Name:     "admins.name",
	SteamID:  "admins.steam_id",
	Immunity: "admins.immunity",
	Flags:    "admins.flags",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var AdminWhere = struct {
	ID       whereHelperint
	Name     whereHelperstring
	SteamID  whereHelperint64
	Immunity whereHelperint
	Flags    whereHelperstring
}{
	ID:       whereHelperint{field: "`admins`.`id`"},
	Name:     whereHelperstring{field: "`admins`.`name`"},
	SteamID:  whereHelperint64{field: "`admins`.`steam_id`"},
	Immunity: whereHelperint{field: "`admins`.`immunity`"},
	Flags:    whereHelperstring{field: "`admins`.`flags`"},
}

// AdminRels is where relationship names are stored.
var AdminRels = struct {
	IssuerAdminBlocks string
}{
	IssuerAdminBlocks: "IssuerAdminBlocks",
}

// adminR is where relationships are stored.
type adminR struct {
	IssuerAdminBlocks BlockSlice `boil:"IssuerAdminBlocks" json:"IssuerAdminBlocks" toml:"IssuerAdminBlocks" yaml:"IssuerAdminBlocks"`
}

// NewStruct creates a new relationship struct
func (*adminR) NewStruct() *adminR {
	return &adminR{}
}

func (r *adminR) GetIssuerAdminBlocks() BlockSlice {
	if r == nil {
		return nil
	}
	return r.IssuerAdminBlocks
}

// adminL is where Load methods for each relationship are stored.
type adminL struct{}

var (
	adminAllColumns            = []string{"id", "name", "steam_id", "immunity", "flags"}
	adminColumnsWithoutDefault = []string{"name", "steam_id", "immunity", "flags"}
	adminColumnsWithDefault    = []string{"id"}
	adminPrimaryKeyColumns     = []string{"id"}
	adminGeneratedColumns      = []string{}
)

type (
	// AdminSlice is an alias for a slice of pointers to Admin.
	// This should almost always be used instead of []Admin.
	AdminSlice []*Admin

	adminQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminType                 = reflect.TypeOf(&Admin{})
	adminMapping              = queries.MakeStructMapping(adminType)
	adminPrimaryKeyMapping, _ = queries.BindMapping(adminType, adminMapping, adminPrimaryKeyColumns)
	adminInsertCacheMut       sync.RWMutex
	adminInsertCache          = make(map[string]insertCache)
	adminUpdateCacheMut       sync.RWMutex
	adminUpdateCache          = make(map[string]updateCache)
	adminUpsertCacheMut       sync.RWMutex
	adminUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single admin record from the query.
func (q adminQuery) One(exec boil.Executor) (*Admin, error) {
	o := &Admin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for admins")
	}

	return o, nil
}

// All returns all Admin records from the query.
func (q adminQuery) All(exec boil.Executor) (AdminSlice, error) {
	var o []*Admin

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Admin slice")
	}

	return o, nil
}

// Count returns the count of all Admin records in the query.
func (q adminQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count admins rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q adminQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if admins exists")
	}

	return count > 0, nil
}

// IssuerAdminBlocks retrieves all the block's Blocks with an executor via issuer_admin_id column.
func (o *Admin) IssuerAdminBlocks(mods ...qm.QueryMod) blockQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`blocks`.`issuer_admin_id`=?", o.ID),
	)

	return Blocks(queryMods...)
}

// LoadIssuerAdminBlocks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (adminL) LoadIssuerAdminBlocks(e boil.Executor, singular bool, maybeAdmin interface{}, mods queries.Applicator) error {
	var slice []*Admin
	var object *Admin

	if singular {
		var ok bool
		object, ok = maybeAdmin.(*Admin)
		if !ok {
			object = new(Admin)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAdmin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAdmin))
			}
		}
	} else {
		s, ok := maybeAdmin.(*[]*Admin)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAdmin)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAdmin))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &adminR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &adminR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`blocks`),
		qm.WhereIn(`blocks.issuer_admin_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load blocks")
	}

	var resultSlice []*Block
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice blocks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on blocks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for blocks")
	}

	if singular {
		object.R.IssuerAdminBlocks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &blockR{}
			}
			foreign.R.IssuerAdmin = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.IssuerAdminID {
				local.R.IssuerAdminBlocks = append(local.R.IssuerAdminBlocks, foreign)
				if foreign.R == nil {
					foreign.R = &blockR{}
				}
				foreign.R.IssuerAdmin = local
				break
			}
		}
	}

	return nil
}

// AddIssuerAdminBlocks adds the given related objects to the existing relationships
// of the admin, optionally inserting them as new records.
// Appends related to o.R.IssuerAdminBlocks.
// Sets related.R.IssuerAdmin appropriately.
func (o *Admin) AddIssuerAdminBlocks(exec boil.Executor, insert bool, related ...*Block) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.IssuerAdminID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `blocks` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"issuer_admin_id"}),
				strmangle.WhereClause("`", "`", 0, blockPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.IssuerAdminID = o.ID
		}
	}

	if o.R == nil {
		o.R = &adminR{
			IssuerAdminBlocks: related,
		}
	} else {
		o.R.IssuerAdminBlocks = append(o.R.IssuerAdminBlocks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &blockR{
				IssuerAdmin: o,
			}
		} else {
			rel.R.IssuerAdmin = o
		}
	}
	return nil
}

// Admins retrieves all the records using an executor.
func Admins(mods ...qm.QueryMod) adminQuery {
	mods = append(mods, qm.From("`admins`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`admins`.*"})
	}

	return adminQuery{q}
}

// FindAdmin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdmin(exec boil.Executor, iD int, selectCols ...string) (*Admin, error) {
	adminObj := &Admin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `admins` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, adminObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from admins")
	}

	return adminObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Admin) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no admins provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(adminColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	adminInsertCacheMut.RLock()
	cache, cached := adminInsertCache[key]
	adminInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			adminAllColumns,
			adminColumnsWithDefault,
			adminColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminType, adminMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `admins` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `admins` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `admins` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, adminPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into admins")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == adminMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for admins")
	}

CacheNoHooks:
	if !cached {
		adminInsertCacheMut.Lock()
		adminInsertCache[key] = cache
		adminInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Admin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Admin) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	adminUpdateCacheMut.RLock()
	cache, cached := adminUpdateCache[key]
	adminUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			adminAllColumns,
			adminPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return errors.New("model: unable to update admins, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `admins` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, adminPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, append(wl, adminPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update admins row")
	}

	if !cached {
		adminUpdateCacheMut.Lock()
		adminUpdateCache[key] = cache
		adminUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q adminQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for admins")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `admins` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in admin slice")
	}

	return nil
}

var mySQLAdminUniqueColumns = []string{
	"id",
	"steam_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Admin) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no admins provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(adminColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAdminUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	adminUpsertCacheMut.RLock()
	cache, cached := adminUpsertCache[key]
	adminUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			adminAllColumns,
			adminColumnsWithDefault,
			adminColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			adminAllColumns,
			adminPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert admins, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`admins`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `admins` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminType, adminMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for admins")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == adminMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(adminType, adminMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for admins")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for admins")
	}

CacheNoHooks:
	if !cached {
		adminUpsertCacheMut.Lock()
		adminUpsertCache[key] = cache
		adminUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Admin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Admin) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Admin provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminPrimaryKeyMapping)
	sql := "DELETE FROM `admins` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from admins")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q adminQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("model: no adminQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from admins")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `admins` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from admin slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Admin) Reload(exec boil.Executor) error {
	ret, err := FindAdmin(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AdminSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `admins`.* FROM `admins` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, adminPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in AdminSlice")
	}

	*o = slice

	return nil
}

// AdminExists checks if the Admin row exists.
func AdminExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `admins` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if admins exists")
	}

	return exists, nil
}

// Exists checks if the Admin row exists.
func (o *Admin) Exists(exec boil.Executor) (bool, error) {
	return AdminExists(exec, o.ID)
}