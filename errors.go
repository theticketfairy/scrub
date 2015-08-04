package scrub

import (
	"fmt"
	"strings"
)

// FieldErrorCode is used to describe the reason for a field validation failure
type FieldErrorCode string

// The standard set of error codes used by the builtin field validators.
const (
	REQUIRED  FieldErrorCode = "required"
	INVALID                  = "invalid"
	MINLENGTH                = "minlength"
	MAXLENGTH                = "maxlength"
	MIN                      = "min"
	MAX                      = "max"
	MULTI                    = "multi"
	CUSTOM                   = "custom"
)

// FieldError is used to capture field validation failure information
type FieldError struct {
	Code   FieldErrorCode `json:"code"`
	Key    string         `json:"key"`
	Detail string         `json:"detail,omitempty"`
	Meta   FieldErrors    `json:"meta,omitempty"`
}

// NewFieldError creates a FieldError for fields
func NewFieldError(code FieldErrorCode, key, detail string) *FieldError {
	return &FieldError{code, key, detail, nil}
}

// NewMultiFieldError creates a FieldError for fields that embed other fields
// such as struct, map and slice fields
func NewMultiFieldError(key string, errs FieldErrors) *FieldError {
	return &FieldError{Key: key, Code: MULTI, Meta: errs}
}

func (ve *FieldError) Error() string {
	out := fmt.Sprintf("[%s] %s", ve.Code, ve.Key)
	if msg := strings.Trim(ve.Detail, " "); msg != "" {
		out = fmt.Sprintf("%s - %s", out, msg)
	}
	return out
}

// FieldErrors is an alias for a slice of field errors pointers
type FieldErrors []*FieldError

func (fe FieldErrors) Error() string {
	out := make([]string, len(fe))
	for i, e := range fe {
		out[i] = e.Error()
	}
	return strings.Join(out, "\n")
}

func describe(fe *FieldError, level int) []string {
	var out []string
	out = append(out, "* "+strings.Repeat("  ", level)+fe.Error())
	for _, e := range fe.Meta {
		n := describe(e, level+1)
		out = append(out, n...)
	}
	return out
}

// Describe recurses through a FieldErrors slice and to build a helpful string
// representation of the validation errors
func (fe FieldErrors) Describe() string {
	var out []string
	for _, e := range fe {
		o := describe(e, 0)
		out = append(out, o...)
	}
	return strings.Join(out, "\n")
}
