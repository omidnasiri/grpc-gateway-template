// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: foo/rpc_foo.proto

package pb

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

// Validate checks the field values on Foo with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Foo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Foo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FooMultiError, or nil if none found.
func (m *Foo) ValidateAll() error {
	return m.validate(true)
}

func (m *Foo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetTitle()) < 3 {
		err := FooValidationError{
			field:  "Title",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FooMultiError(errors)
	}

	return nil
}

// FooMultiError is an error wrapping multiple validation errors returned by
// Foo.ValidateAll() if the designated constraints aren't met.
type FooMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FooMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FooMultiError) AllErrors() []error { return m }

// FooValidationError is the validation error returned by Foo.Validate if the
// designated constraints aren't met.
type FooValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FooValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FooValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FooValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FooValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FooValidationError) ErrorName() string { return "FooValidationError" }

// Error satisfies the builtin error interface
func (e FooValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFoo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FooValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FooValidationError{}
