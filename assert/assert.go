package assert

import (
	"fmt"

	"github.com/patrickhuber/go-types"
)

// True tests the condition. If the condition is false True panics
// True is meant to be used with `defer handle.Error()`
func True(condition bool) {
	Truef(condition, "expected condition to be true")
}

// True tests the condition. If the condition is false True panics
// True is meant to be used with `defer handle.Error()`
func Truef(condition bool, format string, args ...any) {
	if condition {
		return
	}
	types.Throw(fmt.Errorf(format, args...))
}

// False tests the condition. If the condition is true, False panics
// False is meant to be used with `defer handle.Error()`
func False(condition bool) {
	Falsef(condition, "expected condition to be false")
}

// Falsef tests the condition. If the condition is true, Falsef panics
// Falsef is meant to be used with `defer handle.Error()`
func Falsef(condition bool, format string, args ...any) {
	if !condition {
		return
	}
	types.Throw(fmt.Errorf(format, args...))
}

// Nil tests if the value is nil. If the value is not nil, Nil panics.
// Nil is meant to be used with defer handle.Error()
func Nil(value any) {
	Nilf(value, "value is not nil")
}

// NotNil tests if the value is not nil. If the value is nil, NotNil panics.
// NotNil is meant to be used with defer
func NotNil(value any) {
	NotNilf(value, "value is nil")
}

// NotNilf tests if the value is not nil. If the value is nil, NotNilf panics with the given message.
// NotNilf is meant to be used with defer
func NotNilf(value any, format string, args ...any) {
	if value != nil {
		return
	}
	types.Throw(fmt.Errorf(format, args...))
}

// Nilf tests if the value is nil. If the value is not nil, Nilf panics with the given message.
// Nilf is meant to be used with defer handle.Error()
func Nilf(value any, format string, args ...any) {
	if value == nil {
		return
	}
	types.Throw(fmt.Errorf(format, args...))
}
