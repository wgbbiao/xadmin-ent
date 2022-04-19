// Code generated by entc, DO NOT EDIT.

package permission

import (
	"entgo.io/ent/dialect/sql"
	"github.com/wgbbiao/xadminent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ContentTypeID applies equality check predicate on the "content_type_id" field. It's identical to ContentTypeIDEQ.
func ContentTypeID(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentTypeID), v))
	})
}

// ModelCode applies equality check predicate on the "model_code" field. It's identical to ModelCodeEQ.
func ModelCode(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModelCode), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ContentTypeIDEQ applies the EQ predicate on the "content_type_id" field.
func ContentTypeIDEQ(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentTypeID), v))
	})
}

// ContentTypeIDNEQ applies the NEQ predicate on the "content_type_id" field.
func ContentTypeIDNEQ(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentTypeID), v))
	})
}

// ContentTypeIDIn applies the In predicate on the "content_type_id" field.
func ContentTypeIDIn(vs ...uint) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContentTypeID), v...))
	})
}

// ContentTypeIDNotIn applies the NotIn predicate on the "content_type_id" field.
func ContentTypeIDNotIn(vs ...uint) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContentTypeID), v...))
	})
}

// ContentTypeIDGT applies the GT predicate on the "content_type_id" field.
func ContentTypeIDGT(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentTypeID), v))
	})
}

// ContentTypeIDGTE applies the GTE predicate on the "content_type_id" field.
func ContentTypeIDGTE(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentTypeID), v))
	})
}

// ContentTypeIDLT applies the LT predicate on the "content_type_id" field.
func ContentTypeIDLT(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentTypeID), v))
	})
}

// ContentTypeIDLTE applies the LTE predicate on the "content_type_id" field.
func ContentTypeIDLTE(v uint) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentTypeID), v))
	})
}

// ModelCodeEQ applies the EQ predicate on the "model_code" field.
func ModelCodeEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModelCode), v))
	})
}

// ModelCodeNEQ applies the NEQ predicate on the "model_code" field.
func ModelCodeNEQ(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModelCode), v))
	})
}

// ModelCodeIn applies the In predicate on the "model_code" field.
func ModelCodeIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModelCode), v...))
	})
}

// ModelCodeNotIn applies the NotIn predicate on the "model_code" field.
func ModelCodeNotIn(vs ...string) predicate.Permission {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModelCode), v...))
	})
}

// ModelCodeGT applies the GT predicate on the "model_code" field.
func ModelCodeGT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModelCode), v))
	})
}

// ModelCodeGTE applies the GTE predicate on the "model_code" field.
func ModelCodeGTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModelCode), v))
	})
}

// ModelCodeLT applies the LT predicate on the "model_code" field.
func ModelCodeLT(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModelCode), v))
	})
}

// ModelCodeLTE applies the LTE predicate on the "model_code" field.
func ModelCodeLTE(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModelCode), v))
	})
}

// ModelCodeContains applies the Contains predicate on the "model_code" field.
func ModelCodeContains(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModelCode), v))
	})
}

// ModelCodeHasPrefix applies the HasPrefix predicate on the "model_code" field.
func ModelCodeHasPrefix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModelCode), v))
	})
}

// ModelCodeHasSuffix applies the HasSuffix predicate on the "model_code" field.
func ModelCodeHasSuffix(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModelCode), v))
	})
}

// ModelCodeEqualFold applies the EqualFold predicate on the "model_code" field.
func ModelCodeEqualFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModelCode), v))
	})
}

// ModelCodeContainsFold applies the ContainsFold predicate on the "model_code" field.
func ModelCodeContainsFold(v string) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModelCode), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		p(s.Not())
	})
}
