// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-timeline-api/ocp_timeline_api.proto

package ocp_timeline_api

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

// Validate checks the field values on MultiCreateTimelinesV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateTimelinesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetTimelines() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateTimelinesV1RequestValidationError{
					field:  fmt.Sprintf("Timelines[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateTimelinesV1RequestValidationError is the validation error
// returned by MultiCreateTimelinesV1Request.Validate if the designated
// constraints aren't met.
type MultiCreateTimelinesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateTimelinesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateTimelinesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateTimelinesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateTimelinesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateTimelinesV1RequestValidationError) ErrorName() string {
	return "MultiCreateTimelinesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateTimelinesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateTimelinesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateTimelinesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateTimelinesV1RequestValidationError{}

// Validate checks the field values on MultiCreateTimelinesV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateTimelinesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Added

	return nil
}

// MultiCreateTimelinesV1ResponseValidationError is the validation error
// returned by MultiCreateTimelinesV1Response.Validate if the designated
// constraints aren't met.
type MultiCreateTimelinesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateTimelinesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateTimelinesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateTimelinesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateTimelinesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateTimelinesV1ResponseValidationError) ErrorName() string {
	return "MultiCreateTimelinesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateTimelinesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateTimelinesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateTimelinesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateTimelinesV1ResponseValidationError{}

// Validate checks the field values on UpdateTimelineV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateTimelineV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimeline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTimelineV1RequestValidationError{
				field:  "Timeline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateTimelineV1RequestValidationError is the validation error returned by
// UpdateTimelineV1Request.Validate if the designated constraints aren't met.
type UpdateTimelineV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTimelineV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTimelineV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTimelineV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTimelineV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTimelineV1RequestValidationError) ErrorName() string {
	return "UpdateTimelineV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTimelineV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTimelineV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTimelineV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTimelineV1RequestValidationError{}

// Validate checks the field values on UpdateTimelineV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateTimelineV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateTimelineV1ResponseValidationError is the validation error returned by
// UpdateTimelineV1Response.Validate if the designated constraints aren't met.
type UpdateTimelineV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTimelineV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTimelineV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTimelineV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTimelineV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTimelineV1ResponseValidationError) ErrorName() string {
	return "UpdateTimelineV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTimelineV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTimelineV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTimelineV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTimelineV1ResponseValidationError{}

// Validate checks the field values on CreateTimelineV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTimelineV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimeline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTimelineV1RequestValidationError{
				field:  "Timeline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateTimelineV1RequestValidationError is the validation error returned by
// CreateTimelineV1Request.Validate if the designated constraints aren't met.
type CreateTimelineV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTimelineV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTimelineV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTimelineV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTimelineV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTimelineV1RequestValidationError) ErrorName() string {
	return "CreateTimelineV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTimelineV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTimelineV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTimelineV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTimelineV1RequestValidationError{}

// Validate checks the field values on CreateTimelineV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTimelineV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateTimelineV1ResponseValidationError is the validation error returned by
// CreateTimelineV1Response.Validate if the designated constraints aren't met.
type CreateTimelineV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTimelineV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTimelineV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTimelineV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTimelineV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTimelineV1ResponseValidationError) ErrorName() string {
	return "CreateTimelineV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTimelineV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTimelineV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTimelineV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTimelineV1ResponseValidationError{}

// Validate checks the field values on GetTimelineV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTimelineV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetTimelineV1RequestValidationError is the validation error returned by
// GetTimelineV1Request.Validate if the designated constraints aren't met.
type GetTimelineV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimelineV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimelineV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimelineV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimelineV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimelineV1RequestValidationError) ErrorName() string {
	return "GetTimelineV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimelineV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimelineV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimelineV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimelineV1RequestValidationError{}

// Validate checks the field values on GetTimelineV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTimelineV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimeline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTimelineV1ResponseValidationError{
				field:  "Timeline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTimelineV1ResponseValidationError is the validation error returned by
// GetTimelineV1Response.Validate if the designated constraints aren't met.
type GetTimelineV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTimelineV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTimelineV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTimelineV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTimelineV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTimelineV1ResponseValidationError) ErrorName() string {
	return "GetTimelineV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTimelineV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTimelineV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTimelineV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTimelineV1ResponseValidationError{}

// Validate checks the field values on ListTimelineV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListTimelineV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListTimelineV1RequestValidationError is the validation error returned by
// ListTimelineV1Request.Validate if the designated constraints aren't met.
type ListTimelineV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTimelineV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTimelineV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTimelineV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTimelineV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTimelineV1RequestValidationError) ErrorName() string {
	return "ListTimelineV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTimelineV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTimelineV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTimelineV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTimelineV1RequestValidationError{}

// Validate checks the field values on ListTimelineV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListTimelineV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetTimelines() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTimelineV1ResponseValidationError{
					field:  fmt.Sprintf("Timelines[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListTimelineV1ResponseValidationError is the validation error returned by
// ListTimelineV1Response.Validate if the designated constraints aren't met.
type ListTimelineV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTimelineV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTimelineV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTimelineV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTimelineV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTimelineV1ResponseValidationError) ErrorName() string {
	return "ListTimelineV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListTimelineV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTimelineV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTimelineV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTimelineV1ResponseValidationError{}

// Validate checks the field values on RemoveTimelineV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveTimelineV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveTimelineV1RequestValidationError is the validation error returned by
// RemoveTimelineV1Request.Validate if the designated constraints aren't met.
type RemoveTimelineV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTimelineV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTimelineV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTimelineV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTimelineV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTimelineV1RequestValidationError) ErrorName() string {
	return "RemoveTimelineV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTimelineV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTimelineV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTimelineV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTimelineV1RequestValidationError{}

// Validate checks the field values on RemoveTimelineV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveTimelineV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveTimelineV1ResponseValidationError is the validation error returned by
// RemoveTimelineV1Response.Validate if the designated constraints aren't met.
type RemoveTimelineV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveTimelineV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveTimelineV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveTimelineV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveTimelineV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveTimelineV1ResponseValidationError) ErrorName() string {
	return "RemoveTimelineV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveTimelineV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveTimelineV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveTimelineV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveTimelineV1ResponseValidationError{}

// Validate checks the field values on Timeline with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Timeline) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Type

	if v, ok := interface{}(m.GetFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimelineValidationError{
				field:  "From",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TimelineValidationError{
				field:  "To",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TimelineValidationError is the validation error returned by
// Timeline.Validate if the designated constraints aren't met.
type TimelineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimelineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimelineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimelineValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimelineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimelineValidationError) ErrorName() string { return "TimelineValidationError" }

// Error satisfies the builtin error interface
func (e TimelineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeline.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimelineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimelineValidationError{}
