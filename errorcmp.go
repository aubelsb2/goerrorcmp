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
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		if strings.Contains(errStr, expectingErrorOrPrefix) {
			return
		}
		t.Errorf("error %v doesn't contain string: %v", err, expectingErrorOrPrefix)
	case error:
		if errors.Is(err, expectingErrorOrPrefix) {
			return
		}
		t.Errorf("error %v match error: %v", err, expectingErrorOrPrefix)
	default:
		t.Errorf("Unknown type: %v", reflect.TypeOf(expectingErrorOrPrefix))
	}
}
