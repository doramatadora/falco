// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of time.units
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-units/
func Test_Time_units(t *testing.T) {
	tests := []struct {
		unit    string
		time    time.Time
		expect  string
		isError bool
	}{
		{unit: "s", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620s"},
		{unit: "ms", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620000ms"},
		{unit: "us", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620000000us"},
		{unit: "ns", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "1677880620000000000ns"},
		{unit: "z", time: time.Date(2023, 3, 3, 21, 57, 0, 0, time.UTC), expect: "", isError: true},
	}

	for i, tt := range tests {
		ret, err := Time_units(
			&context.Context{},
			&value.String{Value: tt.unit},
			&value.Time{Value: tt.time},
		)
		if err != nil {
			if !tt.isError {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}
			continue
		}
		if ret.Type() != value.StringType {
			t.Errorf("[%d] Unexpected return type, expect=STRING, got=%s", i, ret.Type())
		}
		v := value.Unwrap[*value.String](ret)
		if diff := cmp.Diff(tt.expect, v.Value); diff != "" {
			t.Errorf("[%d] Return value unmatch, diff=%s", i, diff)
		}
	}
}
