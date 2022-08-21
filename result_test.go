package result

import (
	"errors"
	"reflect"
	"testing"
)

func isNil(i any) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if (isNil(expected) && isNil(actual)) || reflect.DeepEqual(expected, actual) {
		return
	}

	t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
}

func TestResult(t *testing.T) {
	{
		r := New(1, nil)
		assertEqual(t, 1, r.Value())
		assertEqual(t, nil, r.Error())
		assertEqual(t, true, r.IsSuccess())
		assertEqual(t, false, r.IsFailure())
		r.Fold(func(value int) {
			assertEqual(t, 1, value)
		}, func(e error) {
			t.Errorf("Expected no error, but got %v", e)
		})
		r.OnSuccess(func(value int) {
			assertEqual(t, 1, value)
		})
		r.OnFailure(func(e error) {
			t.Errorf("Expected no error, but got %v", e)
		})
		r.Recover(func(e error) int {
			t.Errorf("Expected no error, but got %v", e)
			return 0
		})
	}
	{
		v := 1
		r := Success(&v)
		assertEqual(t, 1, *r.Value())
		assertEqual(t, &v, r.Value())
		assertEqual(t, nil, r.Error())
		assertEqual(t, true, r.IsSuccess())
		assertEqual(t, false, r.IsFailure())
		r.Fold(func(value *int) {
			assertEqual(t, 1, *value)
			assertEqual(t, &v, value)
		}, func(e error) {
			t.Errorf("Expected no error, but got %v", e)
		})
		r.OnSuccess(func(value *int) {
			assertEqual(t, 1, *value)
			assertEqual(t, &v, value)
		})
		r.OnFailure(func(e error) {
			t.Errorf("Expected no error, but got %v", e)
		})
		r.Recover(func(e error) *int {
			t.Errorf("Expected no error, but got %v", e)
			return &v
		})
	}
	{
		err := errors.New("error")
		r := Failure[any](err)
		assertEqual(t, nil, r.Value())
		assertEqual(t, err, r.Error())
		assertEqual(t, false, r.IsSuccess())
		assertEqual(t, true, r.IsFailure())
		r.Fold(func(value any) {
			t.Errorf("Expected error, but got value %v", value)
		}, func(e error) {
			assertEqual(t, err, e)
		})
		r.OnSuccess(func(value any) {
			t.Errorf("Expected error, but got value %v", value)
		})
		r.OnFailure(func(e error) {
			assertEqual(t, err, e)
		})
		r.Recover(func(e error) any {
			assertEqual(t, err, e)
			return 1
		})
	}
}
