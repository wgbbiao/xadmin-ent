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
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
	"github.com/wgbbiao/xadminent/ent/xadminrole"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
)

// XadminRoleUpdate is the builder for updating XadminRole entities.
type XadminRoleUpdate struct {
	config
	hooks    []Hook
	mutation *XadminRoleMutation
}

// Where appends a list predicates to the XadminRoleUpdate builder.
func (xru *XadminRoleUpdate) Where(ps ...predicate.XadminRole) *XadminRoleUpdate {
	xru.mutation.Where(ps...)
	return xru
}

// SetName sets the "name" field.
func (xru *XadminRoleUpdate) SetName(s string) *XadminRoleUpdate {
	xru.mutation.SetName(s)
	return xru
}

// SetCreatedAt sets the "created_at" field.
func (xru *XadminRoleUpdate) SetCreatedAt(t time.Time) *XadminRoleUpdate {
	xru.mutation.SetCreatedAt(t)
	return xru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (xru *XadminRoleUpdate) SetNillableCreatedAt(t *time.Time) *XadminRoleUpdate {
	if t != nil {
		xru.SetCreatedAt(*t)
	}
	return xru
}

// SetUpdatedAt sets the "updated_at" field.
func (xru *XadminRoleUpdate) SetUpdatedAt(t time.Time) *XadminRoleUpdate {
	xru.mutation.SetUpdatedAt(t)
	return xru
}

// AddUserIDs adds the "users" edge to the XadminUser entity by IDs.
func (xru *XadminRoleUpdate) AddUserIDs(ids ...int) *XadminRoleUpdate {
	xru.mutation.AddUserIDs(ids...)
	return xru
}

// AddUsers adds the "users" edges to the XadminUser entity.
func (xru *XadminRoleUpdate) AddUsers(x ...*XadminUser) *XadminRoleUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xru.AddUserIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the XadminPermission entity by IDs.
func (xru *XadminRoleUpdate) AddPermissionIDs(ids ...int) *XadminRoleUpdate {
	xru.mutation.AddPermissionIDs(ids...)
	return xru
}

// AddPermissions adds the "permissions" edges to the XadminPermission entity.
func (xru *XadminRoleUpdate) AddPermissions(x ...*XadminPermission) *XadminRoleUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xru.AddPermissionIDs(ids...)
}

// Mutation returns the XadminRoleMutation object of the builder.
func (xru *XadminRoleUpdate) Mutation() *XadminRoleMutation {
	return xru.mutation
}

// ClearUsers clears all "users" edges to the XadminUser entity.
func (xru *XadminRoleUpdate) ClearUsers() *XadminRoleUpdate {
	xru.mutation.ClearUsers()
	return xru
}

// RemoveUserIDs removes the "users" edge to XadminUser entities by IDs.
func (xru *XadminRoleUpdate) RemoveUserIDs(ids ...int) *XadminRoleUpdate {
	xru.mutation.RemoveUserIDs(ids...)
	return xru
}

// RemoveUsers removes "users" edges to XadminUser entities.
func (xru *XadminRoleUpdate) RemoveUsers(x ...*XadminUser) *XadminRoleUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xru.RemoveUserIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the XadminPermission entity.
func (xru *XadminRoleUpdate) ClearPermissions() *XadminRoleUpdate {
	xru.mutation.ClearPermissions()
	return xru
}

// RemovePermissionIDs removes the "permissions" edge to XadminPermission entities by IDs.
func (xru *XadminRoleUpdate) RemovePermissionIDs(ids ...int) *XadminRoleUpdate {
	xru.mutation.RemovePermissionIDs(ids...)
	return xru
}

// RemovePermissions removes "permissions" edges to XadminPermission entities.
func (xru *XadminRoleUpdate) RemovePermissions(x ...*XadminPermission) *XadminRoleUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xru.RemovePermissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xru *XadminRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	xru.defaults()
	if len(xru.hooks) == 0 {
		if err = xru.check(); err != nil {
			return 0, err
		}
		affected, err = xru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xru.check(); err != nil {
				return 0, err
			}
			xru.mutation = mutation
			affected, err = xru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(xru.hooks) - 1; i >= 0; i-- {
			if xru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (xru *XadminRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := xru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xru *XadminRoleUpdate) Exec(ctx context.Context) error {
	_, err := xru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xru *XadminRoleUpdate) ExecX(ctx context.Context) {
	if err := xru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xru *XadminRoleUpdate) defaults() {
	if _, ok := xru.mutation.UpdatedAt(); !ok {
		v := xadminrole.UpdateDefaultUpdatedAt()
		xru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xru *XadminRoleUpdate) check() error {
	if v, ok := xru.mutation.Name(); ok {
		if err := xadminrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "XadminRole.name": %w`, err)}
		}
	}
	return nil
}

func (xru *XadminRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadminrole.Table,
			Columns: xadminrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadminrole.FieldID,
			},
		},
	}
	if ps := xru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminrole.FieldName,
		})
	}
	if value, ok := xru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldCreatedAt,
		})
	}
	if value, ok := xru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldUpdatedAt,
		})
	}
	if xru.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !xru.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xru.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if xru.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xru.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !xru.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xru.mutation.PermissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, xru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadminrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// XadminRoleUpdateOne is the builder for updating a single XadminRole entity.
type XadminRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *XadminRoleMutation
}

// SetName sets the "name" field.
func (xruo *XadminRoleUpdateOne) SetName(s string) *XadminRoleUpdateOne {
	xruo.mutation.SetName(s)
	return xruo
}

// SetCreatedAt sets the "created_at" field.
func (xruo *XadminRoleUpdateOne) SetCreatedAt(t time.Time) *XadminRoleUpdateOne {
	xruo.mutation.SetCreatedAt(t)
	return xruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (xruo *XadminRoleUpdateOne) SetNillableCreatedAt(t *time.Time) *XadminRoleUpdateOne {
	if t != nil {
		xruo.SetCreatedAt(*t)
	}
	return xruo
}

// SetUpdatedAt sets the "updated_at" field.
func (xruo *XadminRoleUpdateOne) SetUpdatedAt(t time.Time) *XadminRoleUpdateOne {
	xruo.mutation.SetUpdatedAt(t)
	return xruo
}

// AddUserIDs adds the "users" edge to the XadminUser entity by IDs.
func (xruo *XadminRoleUpdateOne) AddUserIDs(ids ...int) *XadminRoleUpdateOne {
	xruo.mutation.AddUserIDs(ids...)
	return xruo
}

// AddUsers adds the "users" edges to the XadminUser entity.
func (xruo *XadminRoleUpdateOne) AddUsers(x ...*XadminUser) *XadminRoleUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xruo.AddUserIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the XadminPermission entity by IDs.
func (xruo *XadminRoleUpdateOne) AddPermissionIDs(ids ...int) *XadminRoleUpdateOne {
	xruo.mutation.AddPermissionIDs(ids...)
	return xruo
}

// AddPermissions adds the "permissions" edges to the XadminPermission entity.
func (xruo *XadminRoleUpdateOne) AddPermissions(x ...*XadminPermission) *XadminRoleUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xruo.AddPermissionIDs(ids...)
}

// Mutation returns the XadminRoleMutation object of the builder.
func (xruo *XadminRoleUpdateOne) Mutation() *XadminRoleMutation {
	return xruo.mutation
}

// ClearUsers clears all "users" edges to the XadminUser entity.
func (xruo *XadminRoleUpdateOne) ClearUsers() *XadminRoleUpdateOne {
	xruo.mutation.ClearUsers()
	return xruo
}

// RemoveUserIDs removes the "users" edge to XadminUser entities by IDs.
func (xruo *XadminRoleUpdateOne) RemoveUserIDs(ids ...int) *XadminRoleUpdateOne {
	xruo.mutation.RemoveUserIDs(ids...)
	return xruo
}

// RemoveUsers removes "users" edges to XadminUser entities.
func (xruo *XadminRoleUpdateOne) RemoveUsers(x ...*XadminUser) *XadminRoleUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xruo.RemoveUserIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the XadminPermission entity.
func (xruo *XadminRoleUpdateOne) ClearPermissions() *XadminRoleUpdateOne {
	xruo.mutation.ClearPermissions()
	return xruo
}

// RemovePermissionIDs removes the "permissions" edge to XadminPermission entities by IDs.
func (xruo *XadminRoleUpdateOne) RemovePermissionIDs(ids ...int) *XadminRoleUpdateOne {
	xruo.mutation.RemovePermissionIDs(ids...)
	return xruo
}

// RemovePermissions removes "permissions" edges to XadminPermission entities.
func (xruo *XadminRoleUpdateOne) RemovePermissions(x ...*XadminPermission) *XadminRoleUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xruo.RemovePermissionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xruo *XadminRoleUpdateOne) Select(field string, fields ...string) *XadminRoleUpdateOne {
	xruo.fields = append([]string{field}, fields...)
	return xruo
}

// Save executes the query and returns the updated XadminRole entity.
func (xruo *XadminRoleUpdateOne) Save(ctx context.Context) (*XadminRole, error) {
	var (
		err  error
		node *XadminRole
	)
	xruo.defaults()
	if len(xruo.hooks) == 0 {
		if err = xruo.check(); err != nil {
			return nil, err
		}
		node, err = xruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = xruo.check(); err != nil {
				return nil, err
			}
			xruo.mutation = mutation
			node, err = xruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(xruo.hooks) - 1; i >= 0; i-- {
			if xruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (xruo *XadminRoleUpdateOne) SaveX(ctx context.Context) *XadminRole {
	node, err := xruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xruo *XadminRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := xruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xruo *XadminRoleUpdateOne) ExecX(ctx context.Context) {
	if err := xruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xruo *XadminRoleUpdateOne) defaults() {
	if _, ok := xruo.mutation.UpdatedAt(); !ok {
		v := xadminrole.UpdateDefaultUpdatedAt()
		xruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xruo *XadminRoleUpdateOne) check() error {
	if v, ok := xruo.mutation.Name(); ok {
		if err := xadminrole.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "XadminRole.name": %w`, err)}
		}
	}
	return nil
}

func (xruo *XadminRoleUpdateOne) sqlSave(ctx context.Context) (_node *XadminRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadminrole.Table,
			Columns: xadminrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadminrole.FieldID,
			},
		},
	}
	id, ok := xruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "XadminRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := xruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xadminrole.FieldID)
		for _, f := range fields {
			if !xadminrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != xadminrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminrole.FieldName,
		})
	}
	if value, ok := xruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldCreatedAt,
		})
	}
	if value, ok := xruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminrole.FieldUpdatedAt,
		})
	}
	if xruo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !xruo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xruo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if xruo.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xruo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !xruo.mutation.PermissionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xruo.mutation.PermissionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &XadminRole{config: xruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadminrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
