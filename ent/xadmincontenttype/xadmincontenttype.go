// Code generated by entc, DO NOT EDIT.

package xadmincontenttype

import (
	"time"
)

const (
	// Label holds the string label denoting the xadmincontenttype type in the database.
	Label = "xadmin_contenttype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAppLabel holds the string denoting the app_label field in the database.
	FieldAppLabel = "app_label"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// Table holds the table name of the xadmincontenttype in the database.
	Table = "xadmin_contenttypes"
)

// Columns holds all SQL columns for xadmincontenttype fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAppLabel,
	FieldModel,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// AppLabelValidator is a validator for the "app_label" field. It is called by the builders before save.
	AppLabelValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
)
