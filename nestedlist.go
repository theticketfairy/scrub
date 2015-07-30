package scrub

import "fmt"

// NestedListField is used to capture validation rules for a list of fields
// that hold their own set of validation rules i.e. types implementing Validated
type NestedListField struct {
	Name   string
	Value  []Validated
	Checks []Check
}

// NewNestedListField initialises a NestedListField with a name and value
func NewNestedListField(name string, cast func() []Validated) *NestedListField {
	return &NestedListField{name, cast(), make([]Check, 0)}
}

func (f *NestedListField) check(c Check) {
	f.Checks = append(f.Checks, c)
}

// Custom adds a user defined check for the field value
func (f *NestedListField) Custom(t func(vals []Validated) bool, msg string) {
	f.check(func() (err *FieldError) {
		if !t(f.Value) {
			err = NewFieldError(CUSTOM, f.Name, msg)
		}
		return
	})
}

func (f *NestedListField) validateAll() FieldErrors {
	errs := make(FieldErrors, 0, len(f.Value))
	for i, v := range f.Value {
		name := fmt.Sprintf("%s.%d", f.Name, i)
		if ve := Validate(v); len(ve) > 0 {
			errs = append(errs, NewMultiFieldError(name, ve))
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// Validate iterates (1) over the list of fields' invoking validation on them
// in the same way as a NestedField, then (2) over the field itself running
// any checks and returning the first validation error it encounters
func (f *NestedListField) Validate() (err *FieldError) {
	errs := f.validateAll()
	if len(errs) > 0 {
		err = NewMultiFieldError(f.Name, errs)
	}
	if err != nil {
		return
	}
	for i := 0; i < len(f.Checks) && err == nil; i++ {
		err = f.Checks[i]()
	}
	return
}
