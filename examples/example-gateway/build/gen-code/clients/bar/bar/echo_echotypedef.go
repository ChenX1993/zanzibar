// Code generated by thriftrw v1.8.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// Echo_EchoTypedef_Args represents the arguments for the Echo.echoTypedef function.
//
// The arguments for echoTypedef are sent and received over the wire as this struct.
type Echo_EchoTypedef_Args struct {
	Arg UUID `json:"arg,required"`
}

// ToWire translates a Echo_EchoTypedef_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Echo_EchoTypedef_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = v.Arg.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_EchoTypedef_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoTypedef_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoTypedef_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoTypedef_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Arg, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of Echo_EchoTypedef_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoTypedef_Args
// struct.
func (v *Echo_EchoTypedef_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("Echo_EchoTypedef_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_EchoTypedef_Args match the
// provided Echo_EchoTypedef_Args.
//
// This function performs a deep comparison.
func (v *Echo_EchoTypedef_Args) Equals(rhs *Echo_EchoTypedef_Args) bool {
	if !(v.Arg == rhs.Arg) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoTypedef" for this struct.
func (v *Echo_EchoTypedef_Args) MethodName() string {
	return "echoTypedef"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Echo_EchoTypedef_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Echo_EchoTypedef_Helper provides functions that aid in handling the
// parameters and return values of the Echo.echoTypedef
// function.
var Echo_EchoTypedef_Helper = struct {
	// Args accepts the parameters of echoTypedef in-order and returns
	// the arguments struct for the function.
	Args func(
		arg UUID,
	) *Echo_EchoTypedef_Args

	// IsException returns true if the given error can be thrown
	// by echoTypedef.
	//
	// An error can be thrown by echoTypedef only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoTypedef
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoTypedef into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoTypedef
	//
	//   value, err := echoTypedef(args)
	//   result, err := Echo_EchoTypedef_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoTypedef: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(UUID, error) (*Echo_EchoTypedef_Result, error)

	// UnwrapResponse takes the result struct for echoTypedef
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoTypedef threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Echo_EchoTypedef_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Echo_EchoTypedef_Result) (UUID, error)
}{}

func init() {
	Echo_EchoTypedef_Helper.Args = func(
		arg UUID,
	) *Echo_EchoTypedef_Args {
		return &Echo_EchoTypedef_Args{
			Arg: arg,
		}
	}

	Echo_EchoTypedef_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Echo_EchoTypedef_Helper.WrapResponse = func(success UUID, err error) (*Echo_EchoTypedef_Result, error) {
		if err == nil {
			return &Echo_EchoTypedef_Result{Success: &success}, nil
		}

		return nil, err
	}
	Echo_EchoTypedef_Helper.UnwrapResponse = func(result *Echo_EchoTypedef_Result) (success UUID, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Echo_EchoTypedef_Result represents the result of a Echo.echoTypedef function call.
//
// The result of a echoTypedef execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Echo_EchoTypedef_Result struct {
	// Value returned by echoTypedef after a successful execution.
	Success *UUID `json:"success,omitempty"`
}

// ToWire translates a Echo_EchoTypedef_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Echo_EchoTypedef_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoTypedef_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_EchoTypedef_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoTypedef_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoTypedef_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoTypedef_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x UUID
				x, err = _UUID_Read(field.Value)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Echo_EchoTypedef_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoTypedef_Result
// struct.
func (v *Echo_EchoTypedef_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("Echo_EchoTypedef_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_EchoTypedef_Result match the
// provided Echo_EchoTypedef_Result.
//
// This function performs a deep comparison.
func (v *Echo_EchoTypedef_Result) Equals(rhs *Echo_EchoTypedef_Result) bool {
	if !_UUID_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Echo_EchoTypedef_Result) GetSuccess() (o UUID) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoTypedef" for this struct.
func (v *Echo_EchoTypedef_Result) MethodName() string {
	return "echoTypedef"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Echo_EchoTypedef_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
