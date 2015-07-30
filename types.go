package scrub

// Check is used for capturing a validation rule and returns either a error on
// validation failure or nil
type Check func() *FieldError

// Validated describes types that can build a form holding field validation rules
type Validated interface {
	Form() Form
}

// Field describes types that can run validation
type Field interface {
	Validate() *FieldError
}

// Form is an alias for a slice of fields. i.e. a form is simply a list of fields
type Form []Field

func (f Form) validate() FieldErrors {
	errs := make(FieldErrors, 0, len(f))
	for _, field := range f {
		if err := field.Validate(); err != nil {
			errs = append(errs, field.Validate())
		}
	}
	return errs
}

// Validate uses the form returned by a Validated type and runs validation
// on it return a slice of errors
func Validate(v Validated) FieldErrors {
	return v.Form().validate()
}
