// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
)

// XadminUser is the model entity for the XadminUser schema.
type XadminUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"-"`
	// IsSuper holds the value of the "is_super" field.
	IsSuper bool `json:"is_super"`
	// LastLoginAt holds the value of the "last_login_at" field.
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the XadminUserQuery when eager-loading is set.
	Edges XadminUserEdges `json:"edges"`
}

// XadminUserEdges holds the relations/edges for other nodes in the graph.
type XadminUserEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*XadminRole `json:"roles,omitempty"`
	// Permissions holds the value of the permissions edge.
	Permissions []*XadminPermission `json:"permissions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e XadminUserEdges) RolesOrErr() ([]*XadminRole, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e XadminUserEdges) PermissionsOrErr() ([]*XadminPermission, error) {
	if e.loadedTypes[1] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*XadminUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case xadminuser.FieldIsSuper:
			values[i] = new(sql.NullBool)
		case xadminuser.FieldID:
			values[i] = new(sql.NullInt64)
		case xadminuser.FieldUsername, xadminuser.FieldPassword, xadminuser.FieldSalt:
			values[i] = new(sql.NullString)
		case xadminuser.FieldCreatedAt, xadminuser.FieldUpdatedAt, xadminuser.FieldLastLoginAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type XadminUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the XadminUser fields.
func (xu *XadminUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case xadminuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			xu.ID = int(value.Int64)
		case xadminuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				xu.CreatedAt = value.Time
			}
		case xadminuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				xu.UpdatedAt = value.Time
			}
		case xadminuser.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				xu.Username = value.String
			}
		case xadminuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				xu.Password = value.String
			}
		case xadminuser.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				xu.Salt = value.String
			}
		case xadminuser.FieldIsSuper:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_super", values[i])
			} else if value.Valid {
				xu.IsSuper = value.Bool
			}
		case xadminuser.FieldLastLoginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_at", values[i])
			} else if value.Valid {
				xu.LastLoginAt = new(time.Time)
				*xu.LastLoginAt = value.Time
			}
		}
	}
	return nil
}

// QueryRoles queries the "roles" edge of the XadminUser entity.
func (xu *XadminUser) QueryRoles() *XadminRoleQuery {
	return (&XadminUserClient{config: xu.config}).QueryRoles(xu)
}

// QueryPermissions queries the "permissions" edge of the XadminUser entity.
func (xu *XadminUser) QueryPermissions() *XadminPermissionQuery {
	return (&XadminUserClient{config: xu.config}).QueryPermissions(xu)
}

// Update returns a builder for updating this XadminUser.
// Note that you need to call XadminUser.Unwrap() before calling this method if this XadminUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (xu *XadminUser) Update() *XadminUserUpdateOne {
	return (&XadminUserClient{config: xu.config}).UpdateOne(xu)
}

// Unwrap unwraps the XadminUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (xu *XadminUser) Unwrap() *XadminUser {
	tx, ok := xu.config.driver.(*txDriver)
	if !ok {
		panic("ent: XadminUser is not a transactional entity")
	}
	xu.config.driver = tx.drv
	return xu
}

// String implements the fmt.Stringer.
func (xu *XadminUser) String() string {
	var builder strings.Builder
	builder.WriteString("XadminUser(")
	builder.WriteString(fmt.Sprintf("id=%v", xu.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(xu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(xu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", username=")
	builder.WriteString(xu.Username)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", salt=<sensitive>")
	builder.WriteString(", is_super=")
	builder.WriteString(fmt.Sprintf("%v", xu.IsSuper))
	if v := xu.LastLoginAt; v != nil {
		builder.WriteString(", last_login_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// XadminUsers is a parsable slice of XadminUser.
type XadminUsers []*XadminUser

func (xu XadminUsers) config(cfg config) {
	for _i := range xu {
		xu[_i].config = cfg
	}
}
