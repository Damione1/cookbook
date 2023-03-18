// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProjectSkill is an object representing the database table.
type ProjectSkill struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID int       `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	SkillID   int       `boil:"skill_id" json:"skill_id" toml:"skill_id" yaml:"skill_id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *projectSkillR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectSkillL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectSkillColumns = struct {
	ID        string
	ProjectID string
	SkillID   string
	CreatedAt string
}{
	ID:        "id",
	ProjectID: "project_id",
	SkillID:   "skill_id",
	CreatedAt: "created_at",
}

var ProjectSkillTableColumns = struct {
	ID        string
	ProjectID string
	SkillID   string
	CreatedAt string
}{
	ID:        "project_skills.id",
	ProjectID: "project_skills.project_id",
	SkillID:   "project_skills.skill_id",
	CreatedAt: "project_skills.created_at",
}

// Generated where

var ProjectSkillWhere = struct {
	ID        whereHelperint
	ProjectID whereHelperint
	SkillID   whereHelperint
	CreatedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"project_skills\".\"id\""},
	ProjectID: whereHelperint{field: "\"project_skills\".\"project_id\""},
	SkillID:   whereHelperint{field: "\"project_skills\".\"skill_id\""},
	CreatedAt: whereHelpernull_Time{field: "\"project_skills\".\"created_at\""},
}

// ProjectSkillRels is where relationship names are stored.
var ProjectSkillRels = struct {
	Project string
	Skill   string
}{
	Project: "Project",
	Skill:   "Skill",
}

// projectSkillR is where relationships are stored.
type projectSkillR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
	Skill   *Skill   `boil:"Skill" json:"Skill" toml:"Skill" yaml:"Skill"`
}

// NewStruct creates a new relationship struct
func (*projectSkillR) NewStruct() *projectSkillR {
	return &projectSkillR{}
}

func (r *projectSkillR) GetProject() *Project {
	if r == nil {
		return nil
	}
	return r.Project
}

func (r *projectSkillR) GetSkill() *Skill {
	if r == nil {
		return nil
	}
	return r.Skill
}

// projectSkillL is where Load methods for each relationship are stored.
type projectSkillL struct{}

var (
	projectSkillAllColumns            = []string{"id", "project_id", "skill_id", "created_at"}
	projectSkillColumnsWithoutDefault = []string{"project_id", "skill_id"}
	projectSkillColumnsWithDefault    = []string{"id", "created_at"}
	projectSkillPrimaryKeyColumns     = []string{"id"}
	projectSkillGeneratedColumns      = []string{}
)

type (
	// ProjectSkillSlice is an alias for a slice of pointers to ProjectSkill.
	// This should almost always be used instead of []ProjectSkill.
	ProjectSkillSlice []*ProjectSkill
	// ProjectSkillHook is the signature for custom ProjectSkill hook methods
	ProjectSkillHook func(context.Context, boil.ContextExecutor, *ProjectSkill) error

	projectSkillQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectSkillType                 = reflect.TypeOf(&ProjectSkill{})
	projectSkillMapping              = queries.MakeStructMapping(projectSkillType)
	projectSkillPrimaryKeyMapping, _ = queries.BindMapping(projectSkillType, projectSkillMapping, projectSkillPrimaryKeyColumns)
	projectSkillInsertCacheMut       sync.RWMutex
	projectSkillInsertCache          = make(map[string]insertCache)
	projectSkillUpdateCacheMut       sync.RWMutex
	projectSkillUpdateCache          = make(map[string]updateCache)
	projectSkillUpsertCacheMut       sync.RWMutex
	projectSkillUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectSkillAfterSelectHooks []ProjectSkillHook

var projectSkillBeforeInsertHooks []ProjectSkillHook
var projectSkillAfterInsertHooks []ProjectSkillHook

var projectSkillBeforeUpdateHooks []ProjectSkillHook
var projectSkillAfterUpdateHooks []ProjectSkillHook

var projectSkillBeforeDeleteHooks []ProjectSkillHook
var projectSkillAfterDeleteHooks []ProjectSkillHook

var projectSkillBeforeUpsertHooks []ProjectSkillHook
var projectSkillAfterUpsertHooks []ProjectSkillHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectSkill) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectSkill) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectSkill) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectSkill) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectSkill) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectSkill) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectSkill) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectSkill) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectSkill) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSkillAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectSkillHook registers your hook function for all future operations.
func AddProjectSkillHook(hookPoint boil.HookPoint, projectSkillHook ProjectSkillHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		projectSkillAfterSelectHooks = append(projectSkillAfterSelectHooks, projectSkillHook)
	case boil.BeforeInsertHook:
		projectSkillBeforeInsertHooks = append(projectSkillBeforeInsertHooks, projectSkillHook)
	case boil.AfterInsertHook:
		projectSkillAfterInsertHooks = append(projectSkillAfterInsertHooks, projectSkillHook)
	case boil.BeforeUpdateHook:
		projectSkillBeforeUpdateHooks = append(projectSkillBeforeUpdateHooks, projectSkillHook)
	case boil.AfterUpdateHook:
		projectSkillAfterUpdateHooks = append(projectSkillAfterUpdateHooks, projectSkillHook)
	case boil.BeforeDeleteHook:
		projectSkillBeforeDeleteHooks = append(projectSkillBeforeDeleteHooks, projectSkillHook)
	case boil.AfterDeleteHook:
		projectSkillAfterDeleteHooks = append(projectSkillAfterDeleteHooks, projectSkillHook)
	case boil.BeforeUpsertHook:
		projectSkillBeforeUpsertHooks = append(projectSkillBeforeUpsertHooks, projectSkillHook)
	case boil.AfterUpsertHook:
		projectSkillAfterUpsertHooks = append(projectSkillAfterUpsertHooks, projectSkillHook)
	}
}

// One returns a single projectSkill record from the query.
func (q projectSkillQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectSkill, error) {
	o := &ProjectSkill{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for project_skills")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectSkill records from the query.
func (q projectSkillQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectSkillSlice, error) {
	var o []*ProjectSkill

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProjectSkill slice")
	}

	if len(projectSkillAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectSkill records in the query.
func (q projectSkillQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count project_skills rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectSkillQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if project_skills exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *ProjectSkill) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	return Projects(queryMods...)
}

// Skill pointed to by the foreign key.
func (o *ProjectSkill) Skill(mods ...qm.QueryMod) skillQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SkillID),
	}

	queryMods = append(queryMods, mods...)

	return Skills(queryMods...)
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectSkillL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectSkill interface{}, mods queries.Applicator) error {
	var slice []*ProjectSkill
	var object *ProjectSkill

	if singular {
		var ok bool
		object, ok = maybeProjectSkill.(*ProjectSkill)
		if !ok {
			object = new(ProjectSkill)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProjectSkill)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProjectSkill))
			}
		}
	} else {
		s, ok := maybeProjectSkill.(*[]*ProjectSkill)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProjectSkill)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProjectSkill))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectSkillR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectSkillR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projects`),
		qm.WhereIn(`projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(projectAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.ProjectSkills = append(foreign.R.ProjectSkills, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.ProjectSkills = append(foreign.R.ProjectSkills, local)
				break
			}
		}
	}

	return nil
}

// LoadSkill allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectSkillL) LoadSkill(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectSkill interface{}, mods queries.Applicator) error {
	var slice []*ProjectSkill
	var object *ProjectSkill

	if singular {
		var ok bool
		object, ok = maybeProjectSkill.(*ProjectSkill)
		if !ok {
			object = new(ProjectSkill)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProjectSkill)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProjectSkill))
			}
		}
	} else {
		s, ok := maybeProjectSkill.(*[]*ProjectSkill)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProjectSkill)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProjectSkill))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectSkillR{}
		}
		args = append(args, object.SkillID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectSkillR{}
			}

			for _, a := range args {
				if a == obj.SkillID {
					continue Outer
				}
			}

			args = append(args, obj.SkillID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`skills`),
		qm.WhereIn(`skills.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Skill")
	}

	var resultSlice []*Skill
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Skill")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for skills")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for skills")
	}

	if len(skillAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Skill = foreign
		if foreign.R == nil {
			foreign.R = &skillR{}
		}
		foreign.R.ProjectSkills = append(foreign.R.ProjectSkills, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SkillID == foreign.ID {
				local.R.Skill = foreign
				if foreign.R == nil {
					foreign.R = &skillR{}
				}
				foreign.R.ProjectSkills = append(foreign.R.ProjectSkills, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the projectSkill to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectSkills.
func (o *ProjectSkill) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_skills\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectSkillPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &projectSkillR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			ProjectSkills: ProjectSkillSlice{o},
		}
	} else {
		related.R.ProjectSkills = append(related.R.ProjectSkills, o)
	}

	return nil
}

// SetSkill of the projectSkill to the related item.
// Sets o.R.Skill to related.
// Adds o to related.R.ProjectSkills.
func (o *ProjectSkill) SetSkill(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Skill) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_skills\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"skill_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectSkillPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SkillID = related.ID
	if o.R == nil {
		o.R = &projectSkillR{
			Skill: related,
		}
	} else {
		o.R.Skill = related
	}

	if related.R == nil {
		related.R = &skillR{
			ProjectSkills: ProjectSkillSlice{o},
		}
	} else {
		related.R.ProjectSkills = append(related.R.ProjectSkills, o)
	}

	return nil
}

// ProjectSkills retrieves all the records using an executor.
func ProjectSkills(mods ...qm.QueryMod) projectSkillQuery {
	mods = append(mods, qm.From("\"project_skills\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"project_skills\".*"})
	}

	return projectSkillQuery{q}
}

// FindProjectSkill retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectSkill(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProjectSkill, error) {
	projectSkillObj := &ProjectSkill{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"project_skills\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectSkillObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from project_skills")
	}

	if err = projectSkillObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectSkillObj, err
	}

	return projectSkillObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectSkill) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_skills provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectSkillColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectSkillInsertCacheMut.RLock()
	cache, cached := projectSkillInsertCache[key]
	projectSkillInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectSkillAllColumns,
			projectSkillColumnsWithDefault,
			projectSkillColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectSkillType, projectSkillMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectSkillType, projectSkillMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"project_skills\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"project_skills\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into project_skills")
	}

	if !cached {
		projectSkillInsertCacheMut.Lock()
		projectSkillInsertCache[key] = cache
		projectSkillInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectSkill.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectSkill) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectSkillUpdateCacheMut.RLock()
	cache, cached := projectSkillUpdateCache[key]
	projectSkillUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectSkillAllColumns,
			projectSkillPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update project_skills, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"project_skills\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectSkillPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectSkillType, projectSkillMapping, append(wl, projectSkillPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update project_skills row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for project_skills")
	}

	if !cached {
		projectSkillUpdateCacheMut.Lock()
		projectSkillUpdateCache[key] = cache
		projectSkillUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectSkillQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for project_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for project_skills")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectSkillSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"project_skills\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectSkillPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in projectSkill slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all projectSkill")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectSkill) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_skills provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectSkillColumnsWithDefault, o)

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

	projectSkillUpsertCacheMut.RLock()
	cache, cached := projectSkillUpsertCache[key]
	projectSkillUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectSkillAllColumns,
			projectSkillColumnsWithDefault,
			projectSkillColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			projectSkillAllColumns,
			projectSkillPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert project_skills, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectSkillPrimaryKeyColumns))
			copy(conflict, projectSkillPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"project_skills\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectSkillType, projectSkillMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectSkillType, projectSkillMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert project_skills")
	}

	if !cached {
		projectSkillUpsertCacheMut.Lock()
		projectSkillUpsertCache[key] = cache
		projectSkillUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectSkill record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectSkill) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProjectSkill provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectSkillPrimaryKeyMapping)
	sql := "DELETE FROM \"project_skills\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from project_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for project_skills")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectSkillQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectSkillQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project_skills")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_skills")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectSkillSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectSkillBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"project_skills\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectSkillPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from projectSkill slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_skills")
	}

	if len(projectSkillAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProjectSkill) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectSkill(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectSkillSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectSkillSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSkillPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"project_skills\".* FROM \"project_skills\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectSkillPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectSkillSlice")
	}

	*o = slice

	return nil
}

// ProjectSkillExists checks if the ProjectSkill row exists.
func ProjectSkillExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"project_skills\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if project_skills exists")
	}

	return exists, nil
}

// Exists checks if the ProjectSkill row exists.
func (o *ProjectSkill) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProjectSkillExists(ctx, exec, o.ID)
}
