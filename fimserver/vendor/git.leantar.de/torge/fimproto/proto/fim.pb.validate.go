// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/fim.proto

package proto

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

// Validate checks the field values on StartupInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartupInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartupInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartupInfoMultiError, or
// nil if none found.
func (m *StartupInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *StartupInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CreateBaseline

	// no validation rules for UpdateBaseline

	_StartupInfo_WatchedPaths_Unique := make(map[string]struct{}, len(m.GetWatchedPaths()))

	for idx, item := range m.GetWatchedPaths() {
		_, _ = idx, item

		if _, exists := _StartupInfo_WatchedPaths_Unique[item]; exists {
			err := StartupInfoValidationError{
				field:  fmt.Sprintf("WatchedPaths[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_StartupInfo_WatchedPaths_Unique[item] = struct{}{}
		}

		// no validation rules for WatchedPaths[idx]
	}

	if len(errors) > 0 {
		return StartupInfoMultiError(errors)
	}

	return nil
}

// StartupInfoMultiError is an error wrapping multiple validation errors
// returned by StartupInfo.ValidateAll() if the designated constraints aren't met.
type StartupInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartupInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartupInfoMultiError) AllErrors() []error { return m }

// StartupInfoValidationError is the validation error returned by
// StartupInfo.Validate if the designated constraints aren't met.
type StartupInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartupInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartupInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartupInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartupInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartupInfoValidationError) ErrorName() string { return "StartupInfoValidationError" }

// Error satisfies the builtin error interface
func (e StartupInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartupInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartupInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartupInfoValidationError{}

// Validate checks the field values on FsObject with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FsObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FsObject with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FsObjectMultiError, or nil
// if none found.
func (m *FsObject) ValidateAll() error {
	return m.validate(true)
}

func (m *FsObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPath()) > 1024 {
		err := FsObjectValidationError{
			field:  "Path",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_FsObject_Hash_Pattern.MatchString(m.GetHash()) {
		err := FsObjectValidationError{
			field:  "Hash",
			reason: "value does not match regex pattern \"^$|^[0-9a-fA-F]{64}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Created

	// no validation rules for Modified

	// no validation rules for Uid

	// no validation rules for Gid

	// no validation rules for Mode

	if len(errors) > 0 {
		return FsObjectMultiError(errors)
	}

	return nil
}

// FsObjectMultiError is an error wrapping multiple validation errors returned
// by FsObject.ValidateAll() if the designated constraints aren't met.
type FsObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FsObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FsObjectMultiError) AllErrors() []error { return m }

// FsObjectValidationError is the validation error returned by
// FsObject.Validate if the designated constraints aren't met.
type FsObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FsObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FsObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FsObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FsObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FsObjectValidationError) ErrorName() string { return "FsObjectValidationError" }

// Error satisfies the builtin error interface
func (e FsObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFsObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FsObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FsObjectValidationError{}

var _FsObject_Hash_Pattern = regexp.MustCompile("^$|^[0-9a-fA-F]{64}$")

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Event) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EventMultiError, or nil if none found.
func (m *Event) ValidateAll() error {
	return m.validate(true)
}

func (m *Event) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _Event_Kind_InLookup[m.GetKind()]; !ok {
		err := EventValidationError{
			field:  "Kind",
			reason: "value must be in list [CREATE DELETE CHANGE]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for IssuedAt

	if all {
		switch v := interface{}(m.GetFsObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "FsObject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "FsObject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFsObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "FsObject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EventMultiError(errors)
	}

	return nil
}

// EventMultiError is an error wrapping multiple validation errors returned by
// Event.ValidateAll() if the designated constraints aren't met.
type EventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventMultiError) AllErrors() []error { return m }

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

var _Event_Kind_InLookup = map[string]struct{}{
	"CREATE": {},
	"DELETE": {},
	"CHANGE": {},
}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on Alert with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Alert) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Alert with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AlertMultiError, or nil if none found.
func (m *Alert) ValidateAll() error {
	return m.validate(true)
}

func (m *Alert) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _Alert_Kind_InLookup[m.GetKind()]; !ok {
		err := AlertValidationError{
			field:  "Kind",
			reason: "value must be in list [CREATE DELETE CHANGE]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDifference()) > 1024 {
		err := AlertValidationError{
			field:  "Difference",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPath()) > 1024 {
		err := AlertValidationError{
			field:  "Path",
			reason: "value length must be at most 1024 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for IssuedAt

	if len(errors) > 0 {
		return AlertMultiError(errors)
	}

	return nil
}

// AlertMultiError is an error wrapping multiple validation errors returned by
// Alert.ValidateAll() if the designated constraints aren't met.
type AlertMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AlertMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AlertMultiError) AllErrors() []error { return m }

// AlertValidationError is the validation error returned by Alert.Validate if
// the designated constraints aren't met.
type AlertValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlertValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlertValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlertValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlertValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlertValidationError) ErrorName() string { return "AlertValidationError" }

// Error satisfies the builtin error interface
func (e AlertValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlert.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlertValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlertValidationError{}

var _Alert_Kind_InLookup = map[string]struct{}{
	"CREATE": {},
	"DELETE": {},
	"CHANGE": {},
}

// Validate checks the field values on Agent with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Agent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Agent with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AgentMultiError, or nil if none found.
func (m *Agent) ValidateAll() error {
	return m.validate(true)
}

func (m *Agent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for HasBaseline

	// no validation rules for BaselineIsCurrent

	if len(errors) > 0 {
		return AgentMultiError(errors)
	}

	return nil
}

// AgentMultiError is an error wrapping multiple validation errors returned by
// Agent.ValidateAll() if the designated constraints aren't met.
type AgentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentMultiError) AllErrors() []error { return m }

// AgentValidationError is the validation error returned by Agent.Validate if
// the designated constraints aren't met.
type AgentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentValidationError) ErrorName() string { return "AgentValidationError" }

// Error satisfies the builtin error interface
func (e AgentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentValidationError{}

// Validate checks the field values on EndpointName with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EndpointName) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndpointName with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EndpointNameMultiError, or
// nil if none found.
func (m *EndpointName) ValidateAll() error {
	return m.validate(true)
}

func (m *EndpointName) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) > 32 {
		err := EndpointNameValidationError{
			field:  "Name",
			reason: "value length must be at most 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EndpointNameMultiError(errors)
	}

	return nil
}

// EndpointNameMultiError is an error wrapping multiple validation errors
// returned by EndpointName.ValidateAll() if the designated constraints aren't met.
type EndpointNameMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointNameMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointNameMultiError) AllErrors() []error { return m }

// EndpointNameValidationError is the validation error returned by
// EndpointName.Validate if the designated constraints aren't met.
type EndpointNameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointNameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointNameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointNameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointNameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointNameValidationError) ErrorName() string { return "EndpointNameValidationError" }

// Error satisfies the builtin error interface
func (e EndpointNameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointName.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointNameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointNameValidationError{}

// Validate checks the field values on ClientEndpoint with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientEndpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientEndpoint with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientEndpointMultiError,
// or nil if none found.
func (m *ClientEndpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientEndpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) > 32 {
		err := ClientEndpointValidationError{
			field:  "Name",
			reason: "value length must be at most 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRoles()) < 1 {
		err := ClientEndpointValidationError{
			field:  "Roles",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_ClientEndpoint_Roles_Unique := make(map[string]struct{}, len(m.GetRoles()))

	for idx, item := range m.GetRoles() {
		_, _ = idx, item

		if _, exists := _ClientEndpoint_Roles_Unique[item]; exists {
			err := ClientEndpointValidationError{
				field:  fmt.Sprintf("Roles[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_ClientEndpoint_Roles_Unique[item] = struct{}{}
		}

		if _, ok := _ClientEndpoint_Roles_InLookup[item]; !ok {
			err := ClientEndpointValidationError{
				field:  fmt.Sprintf("Roles[%v]", idx),
				reason: "value must be in list [viewer approver user_admin]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ClientEndpointMultiError(errors)
	}

	return nil
}

// ClientEndpointMultiError is an error wrapping multiple validation errors
// returned by ClientEndpoint.ValidateAll() if the designated constraints
// aren't met.
type ClientEndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientEndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientEndpointMultiError) AllErrors() []error { return m }

// ClientEndpointValidationError is the validation error returned by
// ClientEndpoint.Validate if the designated constraints aren't met.
type ClientEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientEndpointValidationError) ErrorName() string { return "ClientEndpointValidationError" }

// Error satisfies the builtin error interface
func (e ClientEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientEndpointValidationError{}

var _ClientEndpoint_Roles_InLookup = map[string]struct{}{
	"viewer":     {},
	"approver":   {},
	"user_admin": {},
}

// Validate checks the field values on AgentEndpoint with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AgentEndpoint) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AgentEndpoint with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AgentEndpointMultiError, or
// nil if none found.
func (m *AgentEndpoint) ValidateAll() error {
	return m.validate(true)
}

func (m *AgentEndpoint) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) > 32 {
		err := AgentEndpointValidationError{
			field:  "Name",
			reason: "value length must be at most 32 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_AgentEndpoint_WatchedPaths_Unique := make(map[string]struct{}, len(m.GetWatchedPaths()))

	for idx, item := range m.GetWatchedPaths() {
		_, _ = idx, item

		if _, exists := _AgentEndpoint_WatchedPaths_Unique[item]; exists {
			err := AgentEndpointValidationError{
				field:  fmt.Sprintf("WatchedPaths[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_AgentEndpoint_WatchedPaths_Unique[item] = struct{}{}
		}

		// no validation rules for WatchedPaths[idx]
	}

	if len(errors) > 0 {
		return AgentEndpointMultiError(errors)
	}

	return nil
}

// AgentEndpointMultiError is an error wrapping multiple validation errors
// returned by AgentEndpoint.ValidateAll() if the designated constraints
// aren't met.
type AgentEndpointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AgentEndpointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AgentEndpointMultiError) AllErrors() []error { return m }

// AgentEndpointValidationError is the validation error returned by
// AgentEndpoint.Validate if the designated constraints aren't met.
type AgentEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AgentEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AgentEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AgentEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AgentEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AgentEndpointValidationError) ErrorName() string { return "AgentEndpointValidationError" }

// Error satisfies the builtin error interface
func (e AgentEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAgentEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AgentEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AgentEndpointValidationError{}
