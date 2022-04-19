// Code generated by entc, DO NOT EDIT.

package contenttype

import (
	"time"
)

const (
	// Label holds the string label denoting the contenttype type in the database.
	Label = "content_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppLabel holds the string denoting the app_label field in the database.
	FieldAppLabel = "app_label"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the contenttype in the database.
	Table = "content_types"
)

// Columns holds all SQL columns for contenttype fields.
var Columns = []string{
	FieldID,
	FieldAppLabel,
	FieldModel,
	FieldCreatedAt,
	FieldUpdatedAt,
}

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
	// AppLabelValidator is a validator for the "app_label" field. It is called by the builders before save.
	AppLabelValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)