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
		v := 1
		r := Success(&v)
		assertEqual(t, &v, r.Value())
		assertEqual(t, nil, r.Error())
		assertEqual(t, true, r.IsSuccess())
		assertEqual(t, false, r.IsFailure())
	}
	{
		err := errors.New("error")
		r := Failure[any](err)
		assertEqual(t, nil, r.Value())
		assertEqual(t, err, r.Error())
		assertEqual(t, false, r.IsSuccess())
		assertEqual(t, true, r.IsFailure())
	}
}
