package scrub

import "fmt"

// Float64Field is used to capture validation rules for float64 fields
type Float64Field struct {
	Name   string
	Value  float64
	Checks []Check
}

// NewFloat64Field initialises a Float64Field with a name and value
func NewFloat64Field(name string, value float64) *Float64Field {
	var checks []Check
	return &Float64Field{name, value, checks}
}

func (f *Float64Field) check(c Check) {
	f.Checks = append(f.Checks, c)
}

// Min adds a checks for field value being less than a specified minimum
func (f *Float64Field) Min(min float64) {
	msg := fmt.Sprintf("The value of this field must be at least %v", min)
	f.check(func() (err *FieldError) {
		if f.Value < min {
			err = NewFieldError(MIN, f.Name, msg)
		}
		return
	})
}

// Max adds a checks for field value being greater than a specified maximum
func (f *Float64Field) Max(max float64) {
	msg := fmt.Sprintf("The value of this field must be at least %v", max)
	f.check(func() (err *FieldError) {
		if f.Value > max {
			err = NewFieldError(MAX, f.Name, msg)
		}
		return
	})
}

// Between adds a check for field value being within [min, max]
func (f *Float64Field) Between(min, max float64) {
	msg := fmt.Sprintf("The value of this field must be between %v and %v", min, max)
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
func (f *Float64Field) Custom(t func(val float64) bool) {
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
func (f *Float64Field) Validate() (err *FieldError) {
	for i := 0; i < len(f.Checks) && err == nil; i++ {
		err = f.Checks[i]()
	}
	return
}
