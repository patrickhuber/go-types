package assert

import "fmt"

// True tests the condition. If the condition is false True panics
// True is meant to be used with result.Handle()
func True(condition bool) {
	Truef(condition, "expected condition to be true")
}

// True tests the condition. If the condition is false True panics
// True is meant to be used with result.Handle()
func Truef(condition bool, format string, args ...any) {
	if condition {
		return
	}
	panic(fmt.Errorf(format, args...))
}

// False tests the condition. If the condition is true, False panics
// False is meant to be used with result.Handle()
func False(condition bool) {
	Falsef(condition, "expected condition to be false")
}

// Falsef tests the condition. If the condition is true, Falsef panics
// Falsef is meant to be used with result.Handle()
func Falsef(condition bool, format string, args ...any) {
	if !condition {
		return
	}
	panic(fmt.Errorf(format, args...))
}
