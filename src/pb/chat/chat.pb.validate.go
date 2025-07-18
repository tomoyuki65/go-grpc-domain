// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chat/chat.proto

package chat

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

// Validate checks the field values on TextInput with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TextInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TextInput with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TextInputMultiError, or nil
// if none found.
func (m *TextInput) ValidateAll() error {
	return m.validate(true)
}

func (m *TextInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetText()) < 1 {
		err := TextInputValidationError{
			field:  "Text",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TextInputMultiError(errors)
	}

	return nil
}

// TextInputMultiError is an error wrapping multiple validation errors returned
// by TextInput.ValidateAll() if the designated constraints aren't met.
type TextInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TextInputMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TextInputMultiError) AllErrors() []error { return m }

// TextInputValidationError is the validation error returned by
// TextInput.Validate if the designated constraints aren't met.
type TextInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TextInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TextInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TextInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TextInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TextInputValidationError) ErrorName() string { return "TextInputValidationError" }

// Error satisfies the builtin error interface
func (e TextInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTextInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TextInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TextInputValidationError{}

// Validate checks the field values on TextOutput with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TextOutput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TextOutput with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TextOutputMultiError, or
// nil if none found.
func (m *TextOutput) ValidateAll() error {
	return m.validate(true)
}

func (m *TextOutput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Text

	if len(errors) > 0 {
		return TextOutputMultiError(errors)
	}

	return nil
}

// TextOutputMultiError is an error wrapping multiple validation errors
// returned by TextOutput.ValidateAll() if the designated constraints aren't met.
type TextOutputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TextOutputMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TextOutputMultiError) AllErrors() []error { return m }

// TextOutputValidationError is the validation error returned by
// TextOutput.Validate if the designated constraints aren't met.
type TextOutputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TextOutputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TextOutputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TextOutputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TextOutputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TextOutputValidationError) ErrorName() string { return "TextOutputValidationError" }

// Error satisfies the builtin error interface
func (e TextOutputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTextOutput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TextOutputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TextOutputValidationError{}
