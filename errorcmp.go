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
		t.Errorf("error %v doesn't contain string: %v", err, expectingErrorOrPrefix)
	case error:
		if errors.Is(err, expectingErrorOrPrefix) {
			return
		}
		if expectingErrorOrPrefix != nil && ErrorStringContains(err, expectingErrorOrPrefix.Error()) {
			return
		}
		t.Errorf("error %v match error: %v", err, expectingErrorOrPrefix)
	default:
		t.Errorf("Unknown type: %v", reflect.TypeOf(expectingErrorOrPrefix))
	}
}

func ErrorStringContains(err error, str string) bool {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	return strings.Contains(errStr, str)
}
