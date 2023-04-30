// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of std.strpad
// Arguments may be:
// - STRING, INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-strpad/
func Test_Std_strpad(t *testing.T) {
	tests := []struct {
		input  string
		width  int64
		pad    string
		expect string
	}{
		{input: "abc", width: -10, pad: "-=", expect: "abc-=-=-=-"},
		{input: "abc", width: 10, pad: "-=", expect: "-=-=-=-abc"},
		{input: "abcdefghijklmn", width: 10, pad: "-=", expect: "abcdefghijklmn"},
		{input: "abcdefghij", width: 10, pad: "-=", expect: "abcdefghij"},
		{input: "abc", width: 7, pad: "🌸🌼", expect: "🌸abc"},
	}

	for i, tt := range tests {
		ret, err := Std_strpad(
			&context.Context{},
			&value.String{Value: tt.input},
			&value.Integer{Value: tt.width},
			&value.String{Value: tt.pad},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
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
