// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: note.proto

package note_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetNotesListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetNotesListRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetNotesListRequestValidationError is the validation error returned by
// GetNotesListRequest.Validate if the designated constraints aren't met.
type GetNotesListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotesListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotesListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotesListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotesListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotesListRequestValidationError) ErrorName() string {
	return "GetNotesListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotesListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotesListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotesListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotesListRequestValidationError{}

// Validate checks the field values on GetNotesListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetNotesListResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetNotesListResponseValidationError is the validation error returned by
// GetNotesListResponse.Validate if the designated constraints aren't met.
type GetNotesListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNotesListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNotesListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNotesListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNotesListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNotesListResponseValidationError) ErrorName() string {
	return "GetNotesListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetNotesListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNotesListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNotesListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNotesListResponseValidationError{}