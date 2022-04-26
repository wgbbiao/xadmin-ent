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

// XadminUserUpdate is the builder for updating XadminUser entities.
type XadminUserUpdate struct {
	config
	hooks    []Hook
	mutation *XadminUserMutation
}

// Where appends a list predicates to the XadminUserUpdate builder.
func (xuu *XadminUserUpdate) Where(ps ...predicate.XadminUser) *XadminUserUpdate {
	xuu.mutation.Where(ps...)
	return xuu
}

// SetUsername sets the "username" field.
func (xuu *XadminUserUpdate) SetUsername(s string) *XadminUserUpdate {
	xuu.mutation.SetUsername(s)
	return xuu
}

// SetPassword sets the "password" field.
func (xuu *XadminUserUpdate) SetPassword(s string) *XadminUserUpdate {
	xuu.mutation.SetPassword(s)
	return xuu
}

// SetSalt sets the "salt" field.
func (xuu *XadminUserUpdate) SetSalt(s string) *XadminUserUpdate {
	xuu.mutation.SetSalt(s)
	return xuu
}

// SetIsSuper sets the "is_super" field.
func (xuu *XadminUserUpdate) SetIsSuper(b bool) *XadminUserUpdate {
	xuu.mutation.SetIsSuper(b)
	return xuu
}

// SetNillableIsSuper sets the "is_super" field if the given value is not nil.
func (xuu *XadminUserUpdate) SetNillableIsSuper(b *bool) *XadminUserUpdate {
	if b != nil {
		xuu.SetIsSuper(*b)
	}
	return xuu
}

// ClearIsSuper clears the value of the "is_super" field.
func (xuu *XadminUserUpdate) ClearIsSuper() *XadminUserUpdate {
	xuu.mutation.ClearIsSuper()
	return xuu
}

// SetLastLoginAt sets the "last_login_at" field.
func (xuu *XadminUserUpdate) SetLastLoginAt(t time.Time) *XadminUserUpdate {
	xuu.mutation.SetLastLoginAt(t)
	return xuu
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (xuu *XadminUserUpdate) SetNillableLastLoginAt(t *time.Time) *XadminUserUpdate {
	if t != nil {
		xuu.SetLastLoginAt(*t)
	}
	return xuu
}

// ClearLastLoginAt clears the value of the "last_login_at" field.
func (xuu *XadminUserUpdate) ClearLastLoginAt() *XadminUserUpdate {
	xuu.mutation.ClearLastLoginAt()
	return xuu
}

// SetUpdatedAt sets the "updated_at" field.
func (xuu *XadminUserUpdate) SetUpdatedAt(t time.Time) *XadminUserUpdate {
	xuu.mutation.SetUpdatedAt(t)
	return xuu
}

// AddRoleIDs adds the "roles" edge to the XadminRole entity by IDs.
func (xuu *XadminUserUpdate) AddRoleIDs(ids ...int) *XadminUserUpdate {
	xuu.mutation.AddRoleIDs(ids...)
	return xuu
}

// AddRoles adds the "roles" edges to the XadminRole entity.
func (xuu *XadminUserUpdate) AddRoles(x ...*XadminRole) *XadminUserUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuu.AddRoleIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the XadminPermission entity by IDs.
func (xuu *XadminUserUpdate) AddPermissionIDs(ids ...int) *XadminUserUpdate {
	xuu.mutation.AddPermissionIDs(ids...)
	return xuu
}

// AddPermissions adds the "permissions" edges to the XadminPermission entity.
func (xuu *XadminUserUpdate) AddPermissions(x ...*XadminPermission) *XadminUserUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuu.AddPermissionIDs(ids...)
}

// Mutation returns the XadminUserMutation object of the builder.
func (xuu *XadminUserUpdate) Mutation() *XadminUserMutation {
	return xuu.mutation
}

// ClearRoles clears all "roles" edges to the XadminRole entity.
func (xuu *XadminUserUpdate) ClearRoles() *XadminUserUpdate {
	xuu.mutation.ClearRoles()
	return xuu
}

// RemoveRoleIDs removes the "roles" edge to XadminRole entities by IDs.
func (xuu *XadminUserUpdate) RemoveRoleIDs(ids ...int) *XadminUserUpdate {
	xuu.mutation.RemoveRoleIDs(ids...)
	return xuu
}

// RemoveRoles removes "roles" edges to XadminRole entities.
func (xuu *XadminUserUpdate) RemoveRoles(x ...*XadminRole) *XadminUserUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuu.RemoveRoleIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the XadminPermission entity.
func (xuu *XadminUserUpdate) ClearPermissions() *XadminUserUpdate {
	xuu.mutation.ClearPermissions()
	return xuu
}

// RemovePermissionIDs removes the "permissions" edge to XadminPermission entities by IDs.
func (xuu *XadminUserUpdate) RemovePermissionIDs(ids ...int) *XadminUserUpdate {
	xuu.mutation.RemovePermissionIDs(ids...)
	return xuu
}

// RemovePermissions removes "permissions" edges to XadminPermission entities.
func (xuu *XadminUserUpdate) RemovePermissions(x ...*XadminPermission) *XadminUserUpdate {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuu.RemovePermissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (xuu *XadminUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	xuu.defaults()
	if len(xuu.hooks) == 0 {
		affected, err = xuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			xuu.mutation = mutation
			affected, err = xuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(xuu.hooks) - 1; i >= 0; i-- {
			if xuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (xuu *XadminUserUpdate) SaveX(ctx context.Context) int {
	affected, err := xuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (xuu *XadminUserUpdate) Exec(ctx context.Context) error {
	_, err := xuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xuu *XadminUserUpdate) ExecX(ctx context.Context) {
	if err := xuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xuu *XadminUserUpdate) defaults() {
	if _, ok := xuu.mutation.UpdatedAt(); !ok {
		v := xadminuser.UpdateDefaultUpdatedAt()
		xuu.mutation.SetUpdatedAt(v)
	}
}

func (xuu *XadminUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadminuser.Table,
			Columns: xadminuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadminuser.FieldID,
			},
		},
	}
	if ps := xuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xuu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldUsername,
		})
	}
	if value, ok := xuu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldPassword,
		})
	}
	if value, ok := xuu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldSalt,
		})
	}
	if value, ok := xuu.mutation.IsSuper(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: xadminuser.FieldIsSuper,
		})
	}
	if xuu.mutation.IsSuperCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: xadminuser.FieldIsSuper,
		})
	}
	if value, ok := xuu.mutation.LastLoginAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminuser.FieldLastLoginAt,
		})
	}
	if xuu.mutation.LastLoginAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: xadminuser.FieldLastLoginAt,
		})
	}
	if value, ok := xuu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminuser.FieldUpdatedAt,
		})
	}
	if xuu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xuu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !xuu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xuu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if xuu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	if nodes := xuu.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !xuu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	if nodes := xuu.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, xuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// XadminUserUpdateOne is the builder for updating a single XadminUser entity.
type XadminUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *XadminUserMutation
}

// SetUsername sets the "username" field.
func (xuuo *XadminUserUpdateOne) SetUsername(s string) *XadminUserUpdateOne {
	xuuo.mutation.SetUsername(s)
	return xuuo
}

// SetPassword sets the "password" field.
func (xuuo *XadminUserUpdateOne) SetPassword(s string) *XadminUserUpdateOne {
	xuuo.mutation.SetPassword(s)
	return xuuo
}

// SetSalt sets the "salt" field.
func (xuuo *XadminUserUpdateOne) SetSalt(s string) *XadminUserUpdateOne {
	xuuo.mutation.SetSalt(s)
	return xuuo
}

// SetIsSuper sets the "is_super" field.
func (xuuo *XadminUserUpdateOne) SetIsSuper(b bool) *XadminUserUpdateOne {
	xuuo.mutation.SetIsSuper(b)
	return xuuo
}

// SetNillableIsSuper sets the "is_super" field if the given value is not nil.
func (xuuo *XadminUserUpdateOne) SetNillableIsSuper(b *bool) *XadminUserUpdateOne {
	if b != nil {
		xuuo.SetIsSuper(*b)
	}
	return xuuo
}

// ClearIsSuper clears the value of the "is_super" field.
func (xuuo *XadminUserUpdateOne) ClearIsSuper() *XadminUserUpdateOne {
	xuuo.mutation.ClearIsSuper()
	return xuuo
}

// SetLastLoginAt sets the "last_login_at" field.
func (xuuo *XadminUserUpdateOne) SetLastLoginAt(t time.Time) *XadminUserUpdateOne {
	xuuo.mutation.SetLastLoginAt(t)
	return xuuo
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (xuuo *XadminUserUpdateOne) SetNillableLastLoginAt(t *time.Time) *XadminUserUpdateOne {
	if t != nil {
		xuuo.SetLastLoginAt(*t)
	}
	return xuuo
}

// ClearLastLoginAt clears the value of the "last_login_at" field.
func (xuuo *XadminUserUpdateOne) ClearLastLoginAt() *XadminUserUpdateOne {
	xuuo.mutation.ClearLastLoginAt()
	return xuuo
}

// SetUpdatedAt sets the "updated_at" field.
func (xuuo *XadminUserUpdateOne) SetUpdatedAt(t time.Time) *XadminUserUpdateOne {
	xuuo.mutation.SetUpdatedAt(t)
	return xuuo
}

// AddRoleIDs adds the "roles" edge to the XadminRole entity by IDs.
func (xuuo *XadminUserUpdateOne) AddRoleIDs(ids ...int) *XadminUserUpdateOne {
	xuuo.mutation.AddRoleIDs(ids...)
	return xuuo
}

// AddRoles adds the "roles" edges to the XadminRole entity.
func (xuuo *XadminUserUpdateOne) AddRoles(x ...*XadminRole) *XadminUserUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuuo.AddRoleIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the XadminPermission entity by IDs.
func (xuuo *XadminUserUpdateOne) AddPermissionIDs(ids ...int) *XadminUserUpdateOne {
	xuuo.mutation.AddPermissionIDs(ids...)
	return xuuo
}

// AddPermissions adds the "permissions" edges to the XadminPermission entity.
func (xuuo *XadminUserUpdateOne) AddPermissions(x ...*XadminPermission) *XadminUserUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuuo.AddPermissionIDs(ids...)
}

// Mutation returns the XadminUserMutation object of the builder.
func (xuuo *XadminUserUpdateOne) Mutation() *XadminUserMutation {
	return xuuo.mutation
}

// ClearRoles clears all "roles" edges to the XadminRole entity.
func (xuuo *XadminUserUpdateOne) ClearRoles() *XadminUserUpdateOne {
	xuuo.mutation.ClearRoles()
	return xuuo
}

// RemoveRoleIDs removes the "roles" edge to XadminRole entities by IDs.
func (xuuo *XadminUserUpdateOne) RemoveRoleIDs(ids ...int) *XadminUserUpdateOne {
	xuuo.mutation.RemoveRoleIDs(ids...)
	return xuuo
}

// RemoveRoles removes "roles" edges to XadminRole entities.
func (xuuo *XadminUserUpdateOne) RemoveRoles(x ...*XadminRole) *XadminUserUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuuo.RemoveRoleIDs(ids...)
}

// ClearPermissions clears all "permissions" edges to the XadminPermission entity.
func (xuuo *XadminUserUpdateOne) ClearPermissions() *XadminUserUpdateOne {
	xuuo.mutation.ClearPermissions()
	return xuuo
}

// RemovePermissionIDs removes the "permissions" edge to XadminPermission entities by IDs.
func (xuuo *XadminUserUpdateOne) RemovePermissionIDs(ids ...int) *XadminUserUpdateOne {
	xuuo.mutation.RemovePermissionIDs(ids...)
	return xuuo
}

// RemovePermissions removes "permissions" edges to XadminPermission entities.
func (xuuo *XadminUserUpdateOne) RemovePermissions(x ...*XadminPermission) *XadminUserUpdateOne {
	ids := make([]int, len(x))
	for i := range x {
		ids[i] = x[i].ID
	}
	return xuuo.RemovePermissionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (xuuo *XadminUserUpdateOne) Select(field string, fields ...string) *XadminUserUpdateOne {
	xuuo.fields = append([]string{field}, fields...)
	return xuuo
}

// Save executes the query and returns the updated XadminUser entity.
func (xuuo *XadminUserUpdateOne) Save(ctx context.Context) (*XadminUser, error) {
	var (
		err  error
		node *XadminUser
	)
	xuuo.defaults()
	if len(xuuo.hooks) == 0 {
		node, err = xuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*XadminUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			xuuo.mutation = mutation
			node, err = xuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(xuuo.hooks) - 1; i >= 0; i-- {
			if xuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = xuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, xuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (xuuo *XadminUserUpdateOne) SaveX(ctx context.Context) *XadminUser {
	node, err := xuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (xuuo *XadminUserUpdateOne) Exec(ctx context.Context) error {
	_, err := xuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xuuo *XadminUserUpdateOne) ExecX(ctx context.Context) {
	if err := xuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xuuo *XadminUserUpdateOne) defaults() {
	if _, ok := xuuo.mutation.UpdatedAt(); !ok {
		v := xadminuser.UpdateDefaultUpdatedAt()
		xuuo.mutation.SetUpdatedAt(v)
	}
}

func (xuuo *XadminUserUpdateOne) sqlSave(ctx context.Context) (_node *XadminUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   xadminuser.Table,
			Columns: xadminuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: xadminuser.FieldID,
			},
		},
	}
	id, ok := xuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "XadminUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := xuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, xadminuser.FieldID)
		for _, f := range fields {
			if !xadminuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != xadminuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := xuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := xuuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldUsername,
		})
	}
	if value, ok := xuuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldPassword,
		})
	}
	if value, ok := xuuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: xadminuser.FieldSalt,
		})
	}
	if value, ok := xuuo.mutation.IsSuper(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: xadminuser.FieldIsSuper,
		})
	}
	if xuuo.mutation.IsSuperCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: xadminuser.FieldIsSuper,
		})
	}
	if value, ok := xuuo.mutation.LastLoginAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminuser.FieldLastLoginAt,
		})
	}
	if xuuo.mutation.LastLoginAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: xadminuser.FieldLastLoginAt,
		})
	}
	if value, ok := xuuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: xadminuser.FieldUpdatedAt,
		})
	}
	if xuuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xuuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !xuuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := xuuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.RolesTable,
			Columns: xadminuser.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: xadminrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if xuuo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	if nodes := xuuo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !xuuo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	if nodes := xuuo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   xadminuser.PermissionsTable,
			Columns: xadminuser.PermissionsPrimaryKey,
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
	_node = &XadminUser{config: xuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, xuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{xadminuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
