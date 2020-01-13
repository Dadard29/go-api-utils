package test

import "testing"

func AssertError(t *testing.T, msg string, expected interface{}, actual interface{}) {
	t.Errorf("%s | Expected: %v | Actual: %v", msg, expected, actual)
}
