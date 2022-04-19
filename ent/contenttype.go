// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wgbbiao/xadminent/ent/contenttype"
)

// ContentType is the model entity for the ContentType schema.
type ContentType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AppLabel holds the value of the "app_label" field.
	AppLabel string `json:"app_label,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContentType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contenttype.FieldID:
			values[i] = new(sql.NullInt64)
		case contenttype.FieldAppLabel, contenttype.FieldModel:
			values[i] = new(sql.NullString)
		case contenttype.FieldCreatedAt, contenttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ContentType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContentType fields.
func (ct *ContentType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contenttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ct.ID = int(value.Int64)
		case contenttype.FieldAppLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_label", values[i])
			} else if value.Valid {
				ct.AppLabel = value.String
			}
		case contenttype.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				ct.Model = value.String
			}
		case contenttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ct.CreatedAt = value.Time
			}
		case contenttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ct.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ContentType.
// Note that you need to call ContentType.Unwrap() before calling this method if this ContentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ct *ContentType) Update() *ContentTypeUpdateOne {
	return (&ContentTypeClient{config: ct.config}).UpdateOne(ct)
}

// Unwrap unwraps the ContentType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ct *ContentType) Unwrap() *ContentType {
	tx, ok := ct.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContentType is not a transactional entity")
	}
	ct.config.driver = tx.drv
	return ct
}

// String implements the fmt.Stringer.
func (ct *ContentType) String() string {
	var builder strings.Builder
	builder.WriteString("ContentType(")
	builder.WriteString(fmt.Sprintf("id=%v", ct.ID))
	builder.WriteString(", app_label=")
	builder.WriteString(ct.AppLabel)
	builder.WriteString(", model=")
	builder.WriteString(ct.Model)
	builder.WriteString(", created_at=")
	builder.WriteString(ct.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ct.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ContentTypes is a parsable slice of ContentType.
type ContentTypes []*ContentType

func (ct ContentTypes) config(cfg config) {
	for _i := range ct {
		ct[_i].config = cfg
	}
}