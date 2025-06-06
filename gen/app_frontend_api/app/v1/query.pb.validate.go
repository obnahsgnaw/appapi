// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: app_frontend_api/app/v1/query.proto

package appv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DetailRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetailRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetailRequestMultiError, or
// nil if none found.
func (m *DetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DetailRequestMultiError(errors)
	}

	return nil
}

// DetailRequestMultiError is an error wrapping multiple validation errors
// returned by DetailRequest.ValidateAll() if the designated constraints
// aren't met.
type DetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetailRequestMultiError) AllErrors() []error { return m }

// DetailRequestValidationError is the validation error returned by
// DetailRequest.Validate if the designated constraints aren't met.
type DetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailRequestValidationError) ErrorName() string { return "DetailRequestValidationError" }

// Error satisfies the builtin error interface
func (e DetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailRequestValidationError{}

// Validate checks the field values on ValidAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ValidAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidAppRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValidAppRequestMultiError, or nil if none found.
func (m *ValidAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 24 {
		err := ValidAppRequestValidationError{
			field:  "Id",
			reason: "value length must be 24 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	// no validation rules for IgnoreEndtype

	if len(errors) > 0 {
		return ValidAppRequestMultiError(errors)
	}

	return nil
}

// ValidAppRequestMultiError is an error wrapping multiple validation errors
// returned by ValidAppRequest.ValidateAll() if the designated constraints
// aren't met.
type ValidAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidAppRequestMultiError) AllErrors() []error { return m }

// ValidAppRequestValidationError is the validation error returned by
// ValidAppRequest.Validate if the designated constraints aren't met.
type ValidAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidAppRequestValidationError) ErrorName() string { return "ValidAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e ValidAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidAppRequestValidationError{}
