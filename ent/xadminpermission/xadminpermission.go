// Code generated by entc, DO NOT EDIT.

package xadminpermission

import (
	"time"
)

const (
	// Label holds the string label denoting the xadminpermission type in the database.
	Label = "xadmin_permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldXadminPermissionContentType holds the string denoting the xadmin_permission_content_type field in the database.
	FieldXadminPermissionContentType = "xadmin_permission_content_type"
	// EdgeContentType holds the string denoting the contenttype edge name in mutations.
	EdgeContentType = "ContentType"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// Table holds the table name of the xadminpermission in the database.
	Table = "xadmin_permissions"
	// ContentTypeTable is the table that holds the ContentType relation/edge.
	ContentTypeTable = "xadmin_permissions"
	// ContentTypeInverseTable is the table name for the XadminContenttype entity.
	// It exists in this package in order to avoid circular dependency with the "xadmincontenttype" package.
	ContentTypeInverseTable = "xadmin_contenttypes"
	// ContentTypeColumn is the table column denoting the ContentType relation/edge.
	ContentTypeColumn = "xadmin_permission_content_type"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "xadmin_permission_users"
	// UsersInverseTable is the table name for the XadminUser entity.
	// It exists in this package in order to avoid circular dependency with the "xadminuser" package.
	UsersInverseTable = "xadmin_users"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "xadmin_permission_roles"
	// RolesInverseTable is the table name for the XadminRole entity.
	// It exists in this package in order to avoid circular dependency with the "xadminrole" package.
	RolesInverseTable = "xadmin_roles"
)

// Columns holds all SQL columns for xadminpermission fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldXadminPermissionContentType,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"xadmin_permission_id", "xadmin_user_id"}
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"xadmin_permission_id", "xadmin_role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
