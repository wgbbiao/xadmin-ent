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
	"github.com/wgbbiao/xadminent/ent/contenttype"
	"github.com/wgbbiao/xadminent/ent/predicate"
)

// ContentTypeUpdate is the builder for updating ContentType entities.
type ContentTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ContentTypeMutation
}

// Where appends a list predicates to the ContentTypeUpdate builder.
func (ctu *ContentTypeUpdate) Where(ps ...predicate.ContentType) *ContentTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetAppLabel sets the "app_label" field.
func (ctu *ContentTypeUpdate) SetAppLabel(s string) *ContentTypeUpdate {
	ctu.mutation.SetAppLabel(s)
	return ctu
}

// SetModel sets the "model" field.
func (ctu *ContentTypeUpdate) SetModel(s string) *ContentTypeUpdate {
	ctu.mutation.SetModel(s)
	return ctu
}

// SetCreatedAt sets the "created_at" field.
func (ctu *ContentTypeUpdate) SetCreatedAt(t time.Time) *ContentTypeUpdate {
	ctu.mutation.SetCreatedAt(t)
	return ctu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctu *ContentTypeUpdate) SetNillableCreatedAt(t *time.Time) *ContentTypeUpdate {
	if t != nil {
		ctu.SetCreatedAt(*t)
	}
	return ctu
}

// SetUpdatedAt sets the "updated_at" field.
func (ctu *ContentTypeUpdate) SetUpdatedAt(t time.Time) *ContentTypeUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// Mutation returns the ContentTypeMutation object of the builder.
func (ctu *ContentTypeUpdate) Mutation() *ContentTypeMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *ContentTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ctu.defaults()
	if len(ctu.hooks) == 0 {
		if err = ctu.check(); err != nil {
			return 0, err
		}
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctu.check(); err != nil {
				return 0, err
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *ContentTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *ContentTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *ContentTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *ContentTypeUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := contenttype.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *ContentTypeUpdate) check() error {
	if v, ok := ctu.mutation.AppLabel(); ok {
		if err := contenttype.AppLabelValidator(v); err != nil {
			return &ValidationError{Name: "app_label", err: fmt.Errorf(`ent: validator failed for field "ContentType.app_label": %w`, err)}
		}
	}
	if v, ok := ctu.mutation.Model(); ok {
		if err := contenttype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "ContentType.model": %w`, err)}
		}
	}
	return nil
}

func (ctu *ContentTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contenttype.Table,
			Columns: contenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contenttype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.AppLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contenttype.FieldAppLabel,
		})
	}
	if value, ok := ctu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contenttype.FieldModel,
		})
	}
	if value, ok := ctu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contenttype.FieldCreatedAt,
		})
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contenttype.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ContentTypeUpdateOne is the builder for updating a single ContentType entity.
type ContentTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContentTypeMutation
}

// SetAppLabel sets the "app_label" field.
func (ctuo *ContentTypeUpdateOne) SetAppLabel(s string) *ContentTypeUpdateOne {
	ctuo.mutation.SetAppLabel(s)
	return ctuo
}

// SetModel sets the "model" field.
func (ctuo *ContentTypeUpdateOne) SetModel(s string) *ContentTypeUpdateOne {
	ctuo.mutation.SetModel(s)
	return ctuo
}

// SetCreatedAt sets the "created_at" field.
func (ctuo *ContentTypeUpdateOne) SetCreatedAt(t time.Time) *ContentTypeUpdateOne {
	ctuo.mutation.SetCreatedAt(t)
	return ctuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctuo *ContentTypeUpdateOne) SetNillableCreatedAt(t *time.Time) *ContentTypeUpdateOne {
	if t != nil {
		ctuo.SetCreatedAt(*t)
	}
	return ctuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ctuo *ContentTypeUpdateOne) SetUpdatedAt(t time.Time) *ContentTypeUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// Mutation returns the ContentTypeMutation object of the builder.
func (ctuo *ContentTypeUpdateOne) Mutation() *ContentTypeMutation {
	return ctuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *ContentTypeUpdateOne) Select(field string, fields ...string) *ContentTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated ContentType entity.
func (ctuo *ContentTypeUpdateOne) Save(ctx context.Context) (*ContentType, error) {
	var (
		err  error
		node *ContentType
	)
	ctuo.defaults()
	if len(ctuo.hooks) == 0 {
		if err = ctuo.check(); err != nil {
			return nil, err
		}
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctuo.check(); err != nil {
				return nil, err
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *ContentTypeUpdateOne) SaveX(ctx context.Context) *ContentType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *ContentTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *ContentTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *ContentTypeUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := contenttype.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *ContentTypeUpdateOne) check() error {
	if v, ok := ctuo.mutation.AppLabel(); ok {
		if err := contenttype.AppLabelValidator(v); err != nil {
			return &ValidationError{Name: "app_label", err: fmt.Errorf(`ent: validator failed for field "ContentType.app_label": %w`, err)}
		}
	}
	if v, ok := ctuo.mutation.Model(); ok {
		if err := contenttype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "ContentType.model": %w`, err)}
		}
	}
	return nil
}

func (ctuo *ContentTypeUpdateOne) sqlSave(ctx context.Context) (_node *ContentType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contenttype.Table,
			Columns: contenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contenttype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContentType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contenttype.FieldID)
		for _, f := range fields {
			if !contenttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.AppLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contenttype.FieldAppLabel,
		})
	}
	if value, ok := ctuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contenttype.FieldModel,
		})
	}
	if value, ok := ctuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contenttype.FieldCreatedAt,
		})
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contenttype.FieldUpdatedAt,
		})
	}
	_node = &ContentType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}