package scrub

// NestedField is used to capture validation rules for fields that hold their
// own set of validation rules i.e. any type implementing Validated
type NestedField struct {
	Name   string
	Value  Validated
	Checks []Check
}

// NewNestedField initialises a NestedField with a name and value
func NewNestedField(name string, value Validated) *NestedField {
	return &NestedField{name, value, make([]Check, 0)}
}

func (f *NestedField) check(c Check) {
	f.Checks = append(f.Checks, c)
}

// Required adds a check for field value not being nil
func (f *NestedField) Required() {
	f.check(func() (err *FieldError) {
		if f.Value == nil {
			err = NewFieldError(REQUIRED, f.Name, "This field is required")
		}
		return
	})
}

// Custom adds a user defined check for the field value
func (f *NestedField) Custom(t func(val Validated) bool) {
	msg := "The value of this field is invalid"
	f.check(func() (err *FieldError) {
		if !t(f.Value) {
			err = NewFieldError(CUSTOM, f.Name, msg)
		}
		return
	})
}

// Validate iterates (1) over the field's checks and returns the first
// validation error it encounters, then (2) validates the nested fields
// building and returning  multi-field error if any errors were found otherwise
// returns nil
func (f *NestedField) Validate() (err *FieldError) {
	for i := 0; i < len(f.Checks) && err == nil; i++ {
		err = f.Checks[i]()
	}
	if err != nil {
		return
	}
	if errs := Validate(f.Value); len(errs) > 0 {
		err = NewMultiFieldError(f.Name, errs)
	}
	return
}
