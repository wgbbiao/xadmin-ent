// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wgbbiao/xadminent/ent/predicate"
	"github.com/wgbbiao/xadminent/ent/xadmincontenttype"
)

// XadminContenttypeUpdate is the builder for updating XadminContenttype entities.
type XadminContenttypeUpdate struct {
	config
	hooks    []Hook
	mutation *XadminContenttypeMutation
}

// Where appends a list predicates to the XadminContenttypeUpdate builder.
func (xcu *XadminContenttypeUpdate) Where(ps ...predicate.XadminContenttype) *XadminContenttypeUpdate {
	xcu.mutation.Where(ps...)
	return xcu
}

// SetUpdatedAt sets the "updated_at" field.
func (xcu *XadminContenttypeUpdate) SetUpdatedAt(t time.Time) *XadminContenttypeUpdate {
	xcu.mutation.SetUpdatedAt(t)
	return xcu
}

// SetAppLabel sets the "app_label" field.
func (xcu *XadminContenttypeUpdate) SetAppLabel(s string) *XadminContenttypeUpdate {
	xcu.mutation.SetAppLabel(s)
	return xcu
}

// SetModel sets the "model" field.
func (xcu *XadminContenttypeUpdate) SetModel(s string) *XadminContenttypeUpdate {
	xcu.mutation.SetModel(s)
	return xcu
}

// Mutation returns the XadminContenttypeMutation object of the builder.
func (xcu *XadminContenttypeUpdate) Mutation() *XadminContenttypeMutation {
	return xcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xcu *XadminContenttypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	xcu.defaults()
	if len(xcu.hooks) == 0 {
		if err = xcu.check(); err != nil {
			return 0, err
		}
		affected, err = xcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminContenttypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xcu.check(); err != nil {
				return 0, err
			}
			xcu.mutation = mutation
			affected, err = xcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(xcu.hooks) - 1; i >= 0; i-- {
			if xcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (xcu *XadminContenttypeUpdate) SaveX(ctx context.Context) int {
	affected, err := xcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xcu *XadminContenttypeUpdate) Exec(ctx context.Context) error {
	_, err := xcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xcu *XadminContenttypeUpdate) ExecX(ctx context.Context) {
	if err := xcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xcu *XadminContenttypeUpdate) defaults() {
	if _, ok := xcu.mutation.UpdatedAt(); !ok {
		v := xadmincontenttype.UpdateDefaultUpdatedAt()
		xcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xcu *XadminContenttypeUpdate) check() error {
	if v, ok := xcu.mutation.AppLabel(); ok {
		if err := xadmincontenttype.AppLabelValidator(v); err != nil {
			return &ValidationError{Name: "app_label", err: fmt.Errorf(`ent: validator failed for field "XadminContenttype.app_label": %w`, err)}
		}
	}
	if v, ok := xcu.mutation.Model(); ok {
		if err := xadmincontenttype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "XadminContenttype.model": %w`, err)}
		}
	}
	return nil
}

func (xcu *XadminContenttypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadmincontenttype.Table,
			Columns: xadmincontenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadmincontenttype.FieldID,
			},
		},
	}
	if ps := xcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadmincontenttype.FieldUpdatedAt,
		})
	}
	if value, ok := xcu.mutation.AppLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadmincontenttype.FieldAppLabel,
		})
	}
	if value, ok := xcu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadmincontenttype.FieldModel,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, xcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadmincontenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// XadminContenttypeUpdateOne is the builder for updating a single XadminContenttype entity.
type XadminContenttypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *XadminContenttypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (xcuo *XadminContenttypeUpdateOne) SetUpdatedAt(t time.Time) *XadminContenttypeUpdateOne {
	xcuo.mutation.SetUpdatedAt(t)
	return xcuo
}

// SetAppLabel sets the "app_label" field.
func (xcuo *XadminContenttypeUpdateOne) SetAppLabel(s string) *XadminContenttypeUpdateOne {
	xcuo.mutation.SetAppLabel(s)
	return xcuo
}

// SetModel sets the "model" field.
func (xcuo *XadminContenttypeUpdateOne) SetModel(s string) *XadminContenttypeUpdateOne {
	xcuo.mutation.SetModel(s)
	return xcuo
}

// Mutation returns the XadminContenttypeMutation object of the builder.
func (xcuo *XadminContenttypeUpdateOne) Mutation() *XadminContenttypeMutation {
	return xcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xcuo *XadminContenttypeUpdateOne) Select(field string, fields ...string) *XadminContenttypeUpdateOne {
	xcuo.fields = append([]string{field}, fields...)
	return xcuo
}

// Save executes the query and returns the updated XadminContenttype entity.
func (xcuo *XadminContenttypeUpdateOne) Save(ctx context.Context) (*XadminContenttype, error) {
	var (
		err  error
		node *XadminContenttype
	)
	xcuo.defaults()
	if len(xcuo.hooks) == 0 {
		if err = xcuo.check(); err != nil {
			return nil, err
		}
		node, err = xcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminContenttypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xcuo.check(); err != nil {
				return nil, err
			}
			xcuo.mutation = mutation
			node, err = xcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(xcuo.hooks) - 1; i >= 0; i-- {
			if xcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (xcuo *XadminContenttypeUpdateOne) SaveX(ctx context.Context) *XadminContenttype {
	node, err := xcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xcuo *XadminContenttypeUpdateOne) Exec(ctx context.Context) error {
	_, err := xcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xcuo *XadminContenttypeUpdateOne) ExecX(ctx context.Context) {
	if err := xcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xcuo *XadminContenttypeUpdateOne) defaults() {
	if _, ok := xcuo.mutation.UpdatedAt(); !ok {
		v := xadmincontenttype.UpdateDefaultUpdatedAt()
		xcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xcuo *XadminContenttypeUpdateOne) check() error {
	if v, ok := xcuo.mutation.AppLabel(); ok {
		if err := xadmincontenttype.AppLabelValidator(v); err != nil {
			return &ValidationError{Name: "app_label", err: fmt.Errorf(`ent: validator failed for field "XadminContenttype.app_label": %w`, err)}
		}
	}
	if v, ok := xcuo.mutation.Model(); ok {
		if err := xadmincontenttype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "XadminContenttype.model": %w`, err)}
		}
	}
	return nil
}

func (xcuo *XadminContenttypeUpdateOne) sqlSave(ctx context.Context) (_node *XadminContenttype, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadmincontenttype.Table,
			Columns: xadmincontenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadmincontenttype.FieldID,
			},
		},
	}
	id, ok := xcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "XadminContenttype.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := xcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xadmincontenttype.FieldID)
		for _, f := range fields {
			if !xadmincontenttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != xadmincontenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadmincontenttype.FieldUpdatedAt,
		})
	}
	if value, ok := xcuo.mutation.AppLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadmincontenttype.FieldAppLabel,
		})
	}
	if value, ok := xcuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadmincontenttype.FieldModel,
		})
	}
	_node = &XadminContenttype{config: xcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadmincontenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
