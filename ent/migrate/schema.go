// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// XadminContenttypesColumns holds the columns for the "xadmin_contenttypes" table.
	XadminContenttypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "app_label", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// XadminContenttypesTable holds the schema information for the "xadmin_contenttypes" table.
	XadminContenttypesTable = &schema.Table{
		Name:       "xadmin_contenttypes",
		Columns:    XadminContenttypesColumns,
		PrimaryKey: []*schema.Column{XadminContenttypesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "xadmincontenttype_app_label",
				Unique:  false,
				Columns: []*schema.Column{XadminContenttypesColumns[1]},
			},
			{
				Name:    "xadmincontenttype_model",
				Unique:  false,
				Columns: []*schema.Column{XadminContenttypesColumns[2]},
			},
		},
	}
	// XadminPermissionsColumns holds the columns for the "xadmin_permissions" table.
	XadminPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "xadmin_permission_content_type", Type: field.TypeInt, Nullable: true},
	}
	// XadminPermissionsTable holds the schema information for the "xadmin_permissions" table.
	XadminPermissionsTable = &schema.Table{
		Name:       "xadmin_permissions",
		Columns:    XadminPermissionsColumns,
		PrimaryKey: []*schema.Column{XadminPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "xadmin_permissions_xadmin_contenttypes_ContentType",
				Columns:    []*schema.Column{XadminPermissionsColumns[5]},
				RefColumns: []*schema.Column{XadminContenttypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "xadminpermission_name",
				Unique:  false,
				Columns: []*schema.Column{XadminPermissionsColumns[1]},
			},
			{
				Name:    "xadminpermission_code",
				Unique:  false,
				Columns: []*schema.Column{XadminPermissionsColumns[2]},
			},
		},
	}
	// XadminRolesColumns holds the columns for the "xadmin_roles" table.
	XadminRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// XadminRolesTable holds the schema information for the "xadmin_roles" table.
	XadminRolesTable = &schema.Table{
		Name:       "xadmin_roles",
		Columns:    XadminRolesColumns,
		PrimaryKey: []*schema.Column{XadminRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "xadminrole_name",
				Unique:  false,
				Columns: []*schema.Column{XadminRolesColumns[1]},
			},
		},
	}
	// XadminUsersColumns holds the columns for the "xadmin_users" table.
	XadminUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "is_super", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "last_login_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// XadminUsersTable holds the schema information for the "xadmin_users" table.
	XadminUsersTable = &schema.Table{
		Name:       "xadmin_users",
		Columns:    XadminUsersColumns,
		PrimaryKey: []*schema.Column{XadminUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "xadminuser_is_super",
				Unique:  false,
				Columns: []*schema.Column{XadminUsersColumns[4]},
			},
		},
	}
	// XadminPermissionUsersColumns holds the columns for the "xadmin_permission_users" table.
	XadminPermissionUsersColumns = []*schema.Column{
		{Name: "xadmin_permission_id", Type: field.TypeInt},
		{Name: "xadmin_user_id", Type: field.TypeInt},
	}
	// XadminPermissionUsersTable holds the schema information for the "xadmin_permission_users" table.
	XadminPermissionUsersTable = &schema.Table{
		Name:       "xadmin_permission_users",
		Columns:    XadminPermissionUsersColumns,
		PrimaryKey: []*schema.Column{XadminPermissionUsersColumns[0], XadminPermissionUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "xadmin_permission_users_xadmin_permission_id",
				Columns:    []*schema.Column{XadminPermissionUsersColumns[0]},
				RefColumns: []*schema.Column{XadminPermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "xadmin_permission_users_xadmin_user_id",
				Columns:    []*schema.Column{XadminPermissionUsersColumns[1]},
				RefColumns: []*schema.Column{XadminUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// XadminPermissionRolesColumns holds the columns for the "xadmin_permission_roles" table.
	XadminPermissionRolesColumns = []*schema.Column{
		{Name: "xadmin_permission_id", Type: field.TypeInt},
		{Name: "xadmin_role_id", Type: field.TypeInt},
	}
	// XadminPermissionRolesTable holds the schema information for the "xadmin_permission_roles" table.
	XadminPermissionRolesTable = &schema.Table{
		Name:       "xadmin_permission_roles",
		Columns:    XadminPermissionRolesColumns,
		PrimaryKey: []*schema.Column{XadminPermissionRolesColumns[0], XadminPermissionRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "xadmin_permission_roles_xadmin_permission_id",
				Columns:    []*schema.Column{XadminPermissionRolesColumns[0]},
				RefColumns: []*schema.Column{XadminPermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "xadmin_permission_roles_xadmin_role_id",
				Columns:    []*schema.Column{XadminPermissionRolesColumns[1]},
				RefColumns: []*schema.Column{XadminRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// XadminRoleUsersColumns holds the columns for the "xadmin_role_users" table.
	XadminRoleUsersColumns = []*schema.Column{
		{Name: "xadmin_role_id", Type: field.TypeInt},
		{Name: "xadmin_user_id", Type: field.TypeInt},
	}
	// XadminRoleUsersTable holds the schema information for the "xadmin_role_users" table.
	XadminRoleUsersTable = &schema.Table{
		Name:       "xadmin_role_users",
		Columns:    XadminRoleUsersColumns,
		PrimaryKey: []*schema.Column{XadminRoleUsersColumns[0], XadminRoleUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "xadmin_role_users_xadmin_role_id",
				Columns:    []*schema.Column{XadminRoleUsersColumns[0]},
				RefColumns: []*schema.Column{XadminRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "xadmin_role_users_xadmin_user_id",
				Columns:    []*schema.Column{XadminRoleUsersColumns[1]},
				RefColumns: []*schema.Column{XadminUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		XadminContenttypesTable,
		XadminPermissionsTable,
		XadminRolesTable,
		XadminUsersTable,
		XadminPermissionUsersTable,
		XadminPermissionRolesTable,
		XadminRoleUsersTable,
	}
)

func init() {
	XadminPermissionsTable.ForeignKeys[0].RefTable = XadminContenttypesTable
	XadminPermissionUsersTable.ForeignKeys[0].RefTable = XadminPermissionsTable
	XadminPermissionUsersTable.ForeignKeys[1].RefTable = XadminUsersTable
	XadminPermissionRolesTable.ForeignKeys[0].RefTable = XadminPermissionsTable
	XadminPermissionRolesTable.ForeignKeys[1].RefTable = XadminRolesTable
	XadminRoleUsersTable.ForeignKeys[0].RefTable = XadminRolesTable
	XadminRoleUsersTable.ForeignKeys[1].RefTable = XadminUsersTable
}
