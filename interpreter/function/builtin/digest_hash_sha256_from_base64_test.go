// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of digest.hash_sha256_from_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hash-sha256-from-base64/
func Test_Digest_hash_sha256_from_base64(t *testing.T) {
	ret, err := Digest_hash_sha256_from_base64(
		&context.Context{},
		&value.String{Value: "SGVsbG8sIHdvcmxkIQo="},
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret.Type() != value.StringType {
		t.Errorf("Unexpected return type, expect=STRING, got=%s", ret.Type())
	}
	v := value.Unwrap[*value.String](ret)
	expect := "d9014c4624844aa5bac314773d6b689ad467fa4e1d1a50a1b8a99d5a95f72ff5"
	if v.Value != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, v.Value)
	}
}
