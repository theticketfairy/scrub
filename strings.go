package scrub

import (
	"fmt"
	"regexp"
	"strings"
)

// StringField is used to capture validation rules for string fields
type StringField struct {
	Name   string
	Value  string
	Checks []Check
}

// NewStringField initialises a StringField with a name and value
func NewStringField(name, value string) *StringField {
	return &StringField{name, value, make([]Check, 0)}
}

func (f *StringField) check(c Check) {
	f.Checks = append(f.Checks, c)
}

// Required adds a check for field value being an empty string
func (f *StringField) Required() {
	f.check(func() (err *FieldError) {
		if strings.TrimSpace(f.Value) == "" {
			err = NewFieldError(REQUIRED, f.Name, "This field is required")
		}
		return
	})
}

// Matches adds a check for field value matching a regular expression
func (f *StringField) Matches(re *regexp.Regexp) {
	msg := "The value of this field is invalid"
	f.check(func() (err *FieldError) {
		if !re.MatchString(f.Value) {
			err = NewFieldError(INVALID, f.Name, msg)
		}
		return
	})
}

// MinLength adds a check for field value being less than a specified
// character length
func (f *StringField) MinLength(min int) {
	msg := fmt.Sprintf("The value must be at least %d characters long", min)
	f.check(func() (err *FieldError) {
		if len(f.Value) < min {
			err = NewFieldError(MINLENGTH, f.Name, msg)
		}
		return
	})
}

// MaxLength adds a check for field value being more than a specified
// character length
func (f *StringField) MaxLength(max int) {
	msg := fmt.Sprintf("The value must be at most %d characters long", max)
	f.check(func() (err *FieldError) {
		if len(f.Value) > max {
			err = NewFieldError(MAXLENGTH, f.Name, msg)
		}
		return
	})
}

// Custom adds a user defined check for the field value
func (f *StringField) Custom(t func(val string) bool) {
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
func (f *StringField) Validate() (err *FieldError) {
	for i := 0; i < len(f.Checks) && err == nil; i++ {
		err = f.Checks[i]()
	}
	return
}
