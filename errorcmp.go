package goerrorcmp

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func ErrorStringMatchesOrContains(t *testing.T, err error, expectingErrorOrPrefix interface{}) {
	if err == nil && (expectingErrorOrPrefix == nil || expectingErrorOrPrefix == "") {
		return
	}
	if err == expectingErrorOrPrefix {
		return
	}
	switch expectingErrorOrPrefix := expectingErrorOrPrefix.(type) {
	case string:
		if ErrorStringContains(err, expectingErrorOrPrefix) {
			return
		}
		t.Errorf("error `%v` doesn't contain string: `%v`", err, expectingErrorOrPrefix)
	case error:
		if err != nil && errors.Is(err, expectingErrorOrPrefix) {
			return
		}
		if expectingErrorOrPrefix != nil && ErrorStringContains(err, expectingErrorOrPrefix.Error()) {
			return
		}
		t.Errorf("error `%v` match error: `%v`", err, expectingErrorOrPrefix)
	case nil:
		t.Errorf("error `%v` isn't `empty` as expected", err)
	default:
		t.Errorf("Unknown type: %v", reflect.TypeOf(expectingErrorOrPrefix))
	}
}

func ErrorStringContains(err error, str string) bool {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	if str == "" {
		return false
	}
	return strings.Contains(errStr, str)
}
