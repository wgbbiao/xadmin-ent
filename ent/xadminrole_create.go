// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
	"github.com/wgbbiao/xadminent/ent/xadminrole"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
)

// XadminRoleCreate is the builder for creating a XadminRole entity.
type XadminRoleCreate struct {
	config
	mutation *XadminRoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (xrc *XadminRoleCreate) SetCreatedAt(t time.Time) *XadminRoleCreate {
	xrc.mutation.SetCreatedAt(t)
	return xrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (xrc *XadminRoleCreate) SetNillableCreatedAt(t *time.Time) *XadminRoleCreate {
	if t != nil {
		xrc.SetCreatedAt(*t)
	}
	return xrc
}

// SetUpdatedAt sets the "updated_at" field.
func (xrc *XadminRoleCreate) SetUpdatedAt(t time.Time) *XadminRoleCreate {
	xrc.mutation.SetUpdatedAt(t)
	return xrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (xrc *XadminRoleCreate) SetNillableUpdatedAt(t *time.Time) *XadminRoleCreate {
	if t != nil {
		xrc.SetUpdatedAt(*t)
	}
	return xrc
}

// SetName sets the "name" field.
func (xrc *XadminRoleCreate) SetName(s string) *XadminRoleCreate {
	xrc.mutation.SetName(s)
	return xrc
}

// AddUserIDs adds the "users" edge to the XadminUser entity by IDs.
func (xrc *XadminRoleCreate) AddUserIDs(ids ...int) *XadminRoleCreate {
	xrc.mutation.AddUserIDs(ids...)
	return xrc
}

// AddUsers adds the "users" edges to the XadminUser entity.
func (xrc *XadminRoleCreate) AddUsers(x ...*XadminUser) *XadminRoleCreate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xrc.AddUserIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the XadminPermission entity by IDs.
func (xrc *XadminRoleCreate) AddPermissionIDs(ids ...int) *XadminRoleCreate {
	xrc.mutation.AddPermissionIDs(ids...)
	return xrc
}

// AddPermissions adds the "permissions" edges to the XadminPermission entity.
func (xrc *XadminRoleCreate) AddPermissions(x ...*XadminPermission) *XadminRoleCreate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xrc.AddPermissionIDs(ids...)
}

// Mutation returns the XadminRoleMutation object of the builder.
func (xrc *XadminRoleCreate) Mutation() *XadminRoleMutation {
	return xrc.mutation
}

// Save creates the XadminRole in the database.
func (xrc *XadminRoleCreate) Save(ctx context.Context) (*XadminRole, error) {
	var (
		err  error
		node *XadminRole
	)
	xrc.defaults()
	if len(xrc.hooks) == 0 {
		if err = xrc.check(); err != nil {
			return nil, err
		}
		node, err = xrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xrc.check(); err != nil {
				return nil, err
			}
			xrc.mutation = mutation
			if node, err = xrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(xrc.hooks) - 1; i >= 0; i-- {
			if xrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (xrc *XadminRoleCreate) SaveX(ctx context.Context) *XadminRole {
	v, err := xrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (xrc *XadminRoleCreate) Exec(ctx context.Context) error {
	_, err := xrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xrc *XadminRoleCreate) ExecX(ctx context.Context) {
	if err := xrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xrc *XadminRoleCreate) defaults() {
	if _, ok := xrc.mutation.CreatedAt(); !ok {
		v := xadminrole.DefaultCreatedAt()
		xrc.mutation.SetCreatedAt(v)
	}
	if _, ok := xrc.mutation.UpdatedAt(); !ok {
		v := xadminrole.DefaultUpdatedAt()
		xrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xrc *XadminRoleCreate) check() error {
	if _, ok := xrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "XadminRole.created_at"`)}
	}
	if _, ok := xrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "XadminRole.updated_at"`)}
	}
	if _, ok := xrc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "XadminRole.name"`)}
	}
	if v, ok := xrc.mutation.Name(); ok {
		if err := xadminrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "XadminRole.name": %w`, err)}
		}
	}
	return nil
}

func (xrc *XadminRoleCreate) sqlSave(ctx context.Context) (*XadminRole, error) {
	_node, _spec := xrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, xrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (xrc *XadminRoleCreate) createSpec() (*XadminRole, *sqlgraph.CreateSpec) {
	var (
		_node = &XadminRole{config: xrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: xadminrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadminrole.FieldID,
			},
		}
	)
	if value, ok := xrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := xrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := xrc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminrole.FieldName,
		})
		_node.Name = value
	}
	if nodes := xrc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   xadminrole.UsersTable,
			Columns: xadminrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := xrc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminrole.PermissionsTable,
			Columns: xadminrole.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminpermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// XadminRoleCreateBulk is the builder for creating many XadminRole entities in bulk.
type XadminRoleCreateBulk struct {
	config
	builders []*XadminRoleCreate
}

// Save creates the XadminRole entities in the database.
func (xrcb *XadminRoleCreateBulk) Save(ctx context.Context) ([]*XadminRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(xrcb.builders))
	nodes := make([]*XadminRole, len(xrcb.builders))
	mutators := make([]Mutator, len(xrcb.builders))
	for i := range xrcb.builders {
		func(i int, root context.Context) {
			builder := xrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*XadminRoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, xrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, xrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, xrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (xrcb *XadminRoleCreateBulk) SaveX(ctx context.Context) []*XadminRole {
	v, err := xrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (xrcb *XadminRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := xrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xrcb *XadminRoleCreateBulk) ExecX(ctx context.Context) {
	if err := xrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
