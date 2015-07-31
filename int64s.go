package scrub

import "fmt"

// Int64Field is used to capture validation rules for int64 fields
type Int64Field struct {
	Name   string
	Value  int64
	Checks []Check
}

// NewInt64Field initialises a Int64Field with a name and value
func NewInt64Field(name string, value int64) *Int64Field {
	return &Int64Field{name, value, make([]Check, 0)}
}

func (f *Int64Field) check(c Check) {
	f.Checks = append(f.Checks, c)
}

// Min adds a checks for field value being less than a specified minimum
func (f *Int64Field) Min(min int64) {
	msg := fmt.Sprintf("The value of this field must be at least %v", min)
	f.check(func() (err *FieldError) {
		if f.Value < min {
			err = NewFieldError(MIN, f.Name, msg)
		}
		return
	})
}

// Max adds a checks for field value being greater than a specified maximum
func (f *Int64Field) Max(max int64) {
	msg := fmt.Sprintf("The value of this field must be at least %v", max)
	f.check(func() (err *FieldError) {
		if f.Value > max {
			err = NewFieldError(MAX, f.Name, msg)
		}
		return
	})
}

// Between adds a check for field value being within [min, max]
func (f *Int64Field) Between(min, max int64) {
	msg := fmt.Sprintf("The value of this field must be between %d and %d", min, max)
	f.check(func() (err *FieldError) {
		if f.Value < min {
			err = NewFieldError(MIN, f.Name, msg)
		} else if f.Value > max {
			err = NewFieldError(MAX, f.Name, msg)
		}
		return
	})
}

// Custom adds a user defined check for the field value
func (f *Int64Field) Custom(t func(val int64) bool) {
	msg := "The value of this field is invalid"
	f.check(func() (err *FieldError) {
		if !t(f.Value) {
			err = NewFieldError(CUSTOM, f.Name, msg)
		}
		return
	})
}

// Validate iterates over the field's checks and returns the first validation
// error it encounters or nil if no errors found
func (f *Int64Field) Validate() (err *FieldError) {
	for i := 0; i < len(f.Checks) && err == nil; i++ {
		err = f.Checks[i]()
	}
	return
}
