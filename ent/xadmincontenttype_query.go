// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wgbbiao/xadminent/ent/predicate"
	"github.com/wgbbiao/xadminent/ent/xadmincontenttype"
)

// XadminContenttypeQuery is the builder for querying XadminContenttype entities.
type XadminContenttypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.XadminContenttype
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the XadminContenttypeQuery builder.
func (xcq *XadminContenttypeQuery) Where(ps ...predicate.XadminContenttype) *XadminContenttypeQuery {
	xcq.predicates = append(xcq.predicates, ps...)
	return xcq
}

// Limit adds a limit step to the query.
func (xcq *XadminContenttypeQuery) Limit(limit int) *XadminContenttypeQuery {
	xcq.limit = &limit
	return xcq
}

// Offset adds an offset step to the query.
func (xcq *XadminContenttypeQuery) Offset(offset int) *XadminContenttypeQuery {
	xcq.offset = &offset
	return xcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (xcq *XadminContenttypeQuery) Unique(unique bool) *XadminContenttypeQuery {
	xcq.unique = &unique
	return xcq
}

// Order adds an order step to the query.
func (xcq *XadminContenttypeQuery) Order(o ...OrderFunc) *XadminContenttypeQuery {
	xcq.order = append(xcq.order, o...)
	return xcq
}

// First returns the first XadminContenttype entity from the query.
// Returns a *NotFoundError when no XadminContenttype was found.
func (xcq *XadminContenttypeQuery) First(ctx context.Context) (*XadminContenttype, error) {
	nodes, err := xcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{xadmincontenttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) FirstX(ctx context.Context) *XadminContenttype {
	node, err := xcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first XadminContenttype ID from the query.
// Returns a *NotFoundError when no XadminContenttype ID was found.
func (xcq *XadminContenttypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = xcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{xadmincontenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) FirstIDX(ctx context.Context) int {
	id, err := xcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single XadminContenttype entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one XadminContenttype entity is found.
// Returns a *NotFoundError when no XadminContenttype entities are found.
func (xcq *XadminContenttypeQuery) Only(ctx context.Context) (*XadminContenttype, error) {
	nodes, err := xcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{xadmincontenttype.Label}
	default:
		return nil, &NotSingularError{xadmincontenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) OnlyX(ctx context.Context) *XadminContenttype {
	node, err := xcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only XadminContenttype ID in the query.
// Returns a *NotSingularError when more than one XadminContenttype ID is found.
// Returns a *NotFoundError when no entities are found.
func (xcq *XadminContenttypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = xcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = &NotSingularError{xadmincontenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := xcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of XadminContenttypes.
func (xcq *XadminContenttypeQuery) All(ctx context.Context) ([]*XadminContenttype, error) {
	if err := xcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return xcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) AllX(ctx context.Context) []*XadminContenttype {
	nodes, err := xcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of XadminContenttype IDs.
func (xcq *XadminContenttypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := xcq.Select(xadmincontenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) IDsX(ctx context.Context) []int {
	ids, err := xcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (xcq *XadminContenttypeQuery) Count(ctx context.Context) (int, error) {
	if err := xcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return xcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) CountX(ctx context.Context) int {
	count, err := xcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (xcq *XadminContenttypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := xcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return xcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (xcq *XadminContenttypeQuery) ExistX(ctx context.Context) bool {
	exist, err := xcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the XadminContenttypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (xcq *XadminContenttypeQuery) Clone() *XadminContenttypeQuery {
	if xcq == nil {
		return nil
	}
	return &XadminContenttypeQuery{
		config:     xcq.config,
		limit:      xcq.limit,
		offset:     xcq.offset,
		order:      append([]OrderFunc{}, xcq.order...),
		predicates: append([]predicate.XadminContenttype{}, xcq.predicates...),
		// clone intermediate query.
		sql:    xcq.sql.Clone(),
		path:   xcq.path,
		unique: xcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppLabel string `json:"app_label,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.XadminContenttype.Query().
//		GroupBy(xadmincontenttype.FieldAppLabel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (xcq *XadminContenttypeQuery) GroupBy(field string, fields ...string) *XadminContenttypeGroupBy {
	group := &XadminContenttypeGroupBy{config: xcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := xcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return xcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppLabel string `json:"app_label,omitempty"`
//	}
//
//	client.XadminContenttype.Query().
//		Select(xadmincontenttype.FieldAppLabel).
//		Scan(ctx, &v)
//
func (xcq *XadminContenttypeQuery) Select(fields ...string) *XadminContenttypeSelect {
	xcq.fields = append(xcq.fields, fields...)
	return &XadminContenttypeSelect{XadminContenttypeQuery: xcq}
}

func (xcq *XadminContenttypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range xcq.fields {
		if !xadmincontenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if xcq.path != nil {
		prev, err := xcq.path(ctx)
		if err != nil {
			return err
		}
		xcq.sql = prev
	}
	return nil
}

func (xcq *XadminContenttypeQuery) sqlAll(ctx context.Context) ([]*XadminContenttype, error) {
	var (
		nodes = []*XadminContenttype{}
		_spec = xcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &XadminContenttype{config: xcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, xcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (xcq *XadminContenttypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := xcq.querySpec()
	_spec.Node.Columns = xcq.fields
	if len(xcq.fields) > 0 {
		_spec.Unique = xcq.unique != nil && *xcq.unique
	}
	return sqlgraph.CountNodes(ctx, xcq.driver, _spec)
}

func (xcq *XadminContenttypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := xcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (xcq *XadminContenttypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadmincontenttype.Table,
			Columns: xadmincontenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadmincontenttype.FieldID,
			},
		},
		From:   xcq.sql,
		Unique: true,
	}
	if unique := xcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := xcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xadmincontenttype.FieldID)
		for i := range fields {
			if fields[i] != xadmincontenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := xcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := xcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := xcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := xcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (xcq *XadminContenttypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(xcq.driver.Dialect())
	t1 := builder.Table(xadmincontenttype.Table)
	columns := xcq.fields
	if len(columns) == 0 {
		columns = xadmincontenttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if xcq.sql != nil {
		selector = xcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if xcq.unique != nil && *xcq.unique {
		selector.Distinct()
	}
	for _, p := range xcq.predicates {
		p(selector)
	}
	for _, p := range xcq.order {
		p(selector)
	}
	if offset := xcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := xcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// XadminContenttypeGroupBy is the group-by builder for XadminContenttype entities.
type XadminContenttypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (xcgb *XadminContenttypeGroupBy) Aggregate(fns ...AggregateFunc) *XadminContenttypeGroupBy {
	xcgb.fns = append(xcgb.fns, fns...)
	return xcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (xcgb *XadminContenttypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := xcgb.path(ctx)
	if err != nil {
		return err
	}
	xcgb.sql = query
	return xcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := xcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(xcgb.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := xcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := xcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = xcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) StringX(ctx context.Context) string {
	v, err := xcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(xcgb.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := xcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := xcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = xcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) IntX(ctx context.Context) int {
	v, err := xcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(xcgb.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := xcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := xcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = xcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := xcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(xcgb.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := xcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := xcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (xcgb *XadminContenttypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = xcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (xcgb *XadminContenttypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := xcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (xcgb *XadminContenttypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range xcgb.fields {
		if !xadmincontenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := xcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := xcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (xcgb *XadminContenttypeGroupBy) sqlQuery() *sql.Selector {
	selector := xcgb.sql.Select()
	aggregation := make([]string, 0, len(xcgb.fns))
	for _, fn := range xcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(xcgb.fields)+len(xcgb.fns))
		for _, f := range xcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(xcgb.fields...)...)
}

// XadminContenttypeSelect is the builder for selecting fields of XadminContenttype entities.
type XadminContenttypeSelect struct {
	*XadminContenttypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (xcs *XadminContenttypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := xcs.prepareQuery(ctx); err != nil {
		return err
	}
	xcs.sql = xcs.XadminContenttypeQuery.sqlQuery(ctx)
	return xcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := xcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(xcs.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := xcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) StringsX(ctx context.Context) []string {
	v, err := xcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = xcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) StringX(ctx context.Context) string {
	v, err := xcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(xcs.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := xcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) IntsX(ctx context.Context) []int {
	v, err := xcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = xcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) IntX(ctx context.Context) int {
	v, err := xcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(xcs.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := xcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := xcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = xcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) Float64X(ctx context.Context) float64 {
	v, err := xcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(xcs.fields) > 1 {
		return nil, errors.New("ent: XadminContenttypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := xcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := xcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (xcs *XadminContenttypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = xcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{xadmincontenttype.Label}
	default:
		err = fmt.Errorf("ent: XadminContenttypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (xcs *XadminContenttypeSelect) BoolX(ctx context.Context) bool {
	v, err := xcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (xcs *XadminContenttypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := xcs.sql.Query()
	if err := xcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}