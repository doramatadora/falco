// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Std_ip2str_Name = "std.ip2str"

var Std_ip2str_ArgumentTypes = []value.Type{value.IpType}

func Std_ip2str_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Std_ip2str_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Std_ip2str_ArgumentTypes[i] {
			return errors.TypeMismatch(Std_ip2str_Name, i+1, Std_ip2str_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.ip2str
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-ip2str/
func Std_ip2str(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_ip2str_Validate(args); err != nil {
		return value.Null, err
	}

	ip := value.Unwrap[*value.IP](args[0])
	return &value.String{Value: ip.Value.String()}, nil
}
