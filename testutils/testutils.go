package testutils

import (
	"reflect"
	"testing"
)

func ShouldBeEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func ShouldNotBeEqual(t *testing.T, actual, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func ShouldNotBeNil(t *testing.T, v any) {
	if isNil(v) {
		t.Errorf("should be nil, but got: %#v", v)
	}
}

func ShouldBeNil(t *testing.T, v any) {
	if !isNil(v) {
		t.Error("should not be nil")
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	switch value.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:

		return value.IsNil()
	default:
		return false
	}
}
