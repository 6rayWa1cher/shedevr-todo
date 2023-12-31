// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Indicates whether the task is completed or not.
// Ref: #/components/schemas/CompletedEnum
type CompletedEnum string

const (
	CompletedEnumYes       CompletedEnum = "yes"
	CompletedEnumNo        CompletedEnum = "no"
	CompletedEnumCancelled CompletedEnum = "cancelled"
)

// AllValues returns all CompletedEnum values.
func (CompletedEnum) AllValues() []CompletedEnum {
	return []CompletedEnum{
		CompletedEnumYes,
		CompletedEnumNo,
		CompletedEnumCancelled,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CompletedEnum) MarshalText() ([]byte, error) {
	switch s {
	case CompletedEnumYes:
		return []byte(s), nil
	case CompletedEnumNo:
		return []byte(s), nil
	case CompletedEnumCancelled:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CompletedEnum) UnmarshalText(data []byte) error {
	switch CompletedEnum(data) {
	case CompletedEnumYes:
		*s = CompletedEnumYes
		return nil
	case CompletedEnumNo:
		*s = CompletedEnumNo
		return nil
	case CompletedEnumCancelled:
		*s = CompletedEnumCancelled
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Counter
type Counter struct {
	// The value of the counter.
	Value float64 `json:"value"`
	// The scale of the counter.
	Scale string `json:"scale"`
	// The max value of the counter.
	MaxValue float64 `json:"max_value"`
}

// GetValue returns the value of Value.
func (s *Counter) GetValue() float64 {
	return s.Value
}

// GetScale returns the value of Scale.
func (s *Counter) GetScale() string {
	return s.Scale
}

// GetMaxValue returns the value of MaxValue.
func (s *Counter) GetMaxValue() float64 {
	return s.MaxValue
}

// SetValue sets the value of Value.
func (s *Counter) SetValue(val float64) {
	s.Value = val
}

// SetScale sets the value of Scale.
func (s *Counter) SetScale(val string) {
	s.Scale = val
}

// SetMaxValue sets the value of MaxValue.
func (s *Counter) SetMaxValue(val float64) {
	s.MaxValue = val
}

// DeleteTaskNoContent is response for DeleteTask operation.
type DeleteTaskNoContent struct{}

func (*DeleteTaskNoContent) deleteTaskRes() {}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/NewTask
type NewTask struct {
	// The title of the new task.
	Title string `json:"title"`
	// The description of the task.
	Text      OptString        `json:"text"`
	Completed OptCompletedEnum `json:"completed"`
	Counter   OptCounter       `json:"counter"`
}

// GetTitle returns the value of Title.
func (s *NewTask) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *NewTask) GetText() OptString {
	return s.Text
}

// GetCompleted returns the value of Completed.
func (s *NewTask) GetCompleted() OptCompletedEnum {
	return s.Completed
}

// GetCounter returns the value of Counter.
func (s *NewTask) GetCounter() OptCounter {
	return s.Counter
}

// SetTitle sets the value of Title.
func (s *NewTask) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *NewTask) SetText(val OptString) {
	s.Text = val
}

// SetCompleted sets the value of Completed.
func (s *NewTask) SetCompleted(val OptCompletedEnum) {
	s.Completed = val
}

// SetCounter sets the value of Counter.
func (s *NewTask) SetCounter(val OptCounter) {
	s.Counter = val
}

// NewOptCompletedEnum returns new OptCompletedEnum with value set to v.
func NewOptCompletedEnum(v CompletedEnum) OptCompletedEnum {
	return OptCompletedEnum{
		Value: v,
		Set:   true,
	}
}

// OptCompletedEnum is optional CompletedEnum.
type OptCompletedEnum struct {
	Value CompletedEnum
	Set   bool
}

// IsSet returns true if OptCompletedEnum was set.
func (o OptCompletedEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCompletedEnum) Reset() {
	var v CompletedEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCompletedEnum) SetTo(v CompletedEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCompletedEnum) Get() (v CompletedEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCompletedEnum) Or(d CompletedEnum) CompletedEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptCounter returns new OptCounter with value set to v.
func NewOptCounter(v Counter) OptCounter {
	return OptCounter{
		Value: v,
		Set:   true,
	}
}

// OptCounter is optional Counter.
type OptCounter struct {
	Value Counter
	Set   bool
}

// IsSet returns true if OptCounter was set.
func (o OptCounter) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptCounter) Reset() {
	var v Counter
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptCounter) SetTo(v Counter) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptCounter) Get() (v Counter, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptCounter) Or(d Counter) Counter {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type RemoteUserAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *RemoteUserAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *RemoteUserAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// Ref: #/components/schemas/Task
type Task struct {
	// The unique identifier for the task.
	ID OptInt64 `json:"id"`
	// The title of the task.
	Title string `json:"title"`
	// The description of the task.
	Text      string        `json:"text"`
	Completed CompletedEnum `json:"completed"`
	Counter   OptCounter    `json:"counter"`
}

// GetID returns the value of ID.
func (s *Task) GetID() OptInt64 {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Task) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *Task) GetText() string {
	return s.Text
}

// GetCompleted returns the value of Completed.
func (s *Task) GetCompleted() CompletedEnum {
	return s.Completed
}

// GetCounter returns the value of Counter.
func (s *Task) GetCounter() OptCounter {
	return s.Counter
}

// SetID sets the value of ID.
func (s *Task) SetID(val OptInt64) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Task) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *Task) SetText(val string) {
	s.Text = val
}

// SetCompleted sets the value of Completed.
func (s *Task) SetCompleted(val CompletedEnum) {
	s.Completed = val
}

// SetCounter sets the value of Counter.
func (s *Task) SetCounter(val OptCounter) {
	s.Counter = val
}

func (*Task) getTaskByIdRes() {}
func (*Task) updateTaskRes()  {}

type TaskNotFound struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *TaskNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *TaskNotFound) SetError(val OptString) {
	s.Error = val
}

func (*TaskNotFound) deleteTaskRes()  {}
func (*TaskNotFound) getTaskByIdRes() {}
func (*TaskNotFound) updateTaskRes()  {}

// Ref: #/components/schemas/UpdateTask
type UpdateTask struct {
	// The unique identifier for the task.
	ID OptInt64 `json:"id"`
	// The title of the new task.
	Title string `json:"title"`
	// The description of the task.
	Text      OptString        `json:"text"`
	Completed OptCompletedEnum `json:"completed"`
	Counter   OptCounter       `json:"counter"`
}

// GetID returns the value of ID.
func (s *UpdateTask) GetID() OptInt64 {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *UpdateTask) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *UpdateTask) GetText() OptString {
	return s.Text
}

// GetCompleted returns the value of Completed.
func (s *UpdateTask) GetCompleted() OptCompletedEnum {
	return s.Completed
}

// GetCounter returns the value of Counter.
func (s *UpdateTask) GetCounter() OptCounter {
	return s.Counter
}

// SetID sets the value of ID.
func (s *UpdateTask) SetID(val OptInt64) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *UpdateTask) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *UpdateTask) SetText(val OptString) {
	s.Text = val
}

// SetCompleted sets the value of Completed.
func (s *UpdateTask) SetCompleted(val OptCompletedEnum) {
	s.Completed = val
}

// SetCounter sets the value of Counter.
func (s *UpdateTask) SetCounter(val OptCounter) {
	s.Counter = val
}
