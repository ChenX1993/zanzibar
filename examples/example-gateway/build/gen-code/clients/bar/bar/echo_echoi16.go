// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Echo_EchoI16_Args represents the arguments for the Echo.echoI16 function.
//
// The arguments for echoI16 are sent and received over the wire as this struct.
type Echo_EchoI16_Args struct {
	Arg int16 `json:"arg,required"`
}

// ToWire translates a Echo_EchoI16_Args struct into a Thrift-level intermediate
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
func (v *Echo_EchoI16_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueI16(v.Arg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_EchoI16_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoI16_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoI16_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoI16_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI16 {
				v.Arg, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of Echo_EchoI16_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoI16_Args
// struct.
func (v *Echo_EchoI16_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("Echo_EchoI16_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_EchoI16_Args match the
// provided Echo_EchoI16_Args.
//
// This function performs a deep comparison.
func (v *Echo_EchoI16_Args) Equals(rhs *Echo_EchoI16_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Arg == rhs.Arg) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_EchoI16_Args.
func (v *Echo_EchoI16_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddInt16("arg", v.Arg)
	return err
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *Echo_EchoI16_Args) GetArg() (o int16) { return v.Arg }

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoI16" for this struct.
func (v *Echo_EchoI16_Args) MethodName() string {
	return "echoI16"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Echo_EchoI16_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Echo_EchoI16_Helper provides functions that aid in handling the
// parameters and return values of the Echo.echoI16
// function.
var Echo_EchoI16_Helper = struct {
	// Args accepts the parameters of echoI16 in-order and returns
	// the arguments struct for the function.
	Args func(
		arg int16,
	) *Echo_EchoI16_Args

	// IsException returns true if the given error can be thrown
	// by echoI16.
	//
	// An error can be thrown by echoI16 only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoI16
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoI16 into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoI16
	//
	//   value, err := echoI16(args)
	//   result, err := Echo_EchoI16_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoI16: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(int16, error) (*Echo_EchoI16_Result, error)

	// UnwrapResponse takes the result struct for echoI16
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoI16 threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Echo_EchoI16_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Echo_EchoI16_Result) (int16, error)
}{}

func init() {
	Echo_EchoI16_Helper.Args = func(
		arg int16,
	) *Echo_EchoI16_Args {
		return &Echo_EchoI16_Args{
			Arg: arg,
		}
	}

	Echo_EchoI16_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Echo_EchoI16_Helper.WrapResponse = func(success int16, err error) (*Echo_EchoI16_Result, error) {
		if err == nil {
			return &Echo_EchoI16_Result{Success: &success}, nil
		}

		return nil, err
	}
	Echo_EchoI16_Helper.UnwrapResponse = func(result *Echo_EchoI16_Result) (success int16, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Echo_EchoI16_Result represents the result of a Echo.echoI16 function call.
//
// The result of a echoI16 execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Echo_EchoI16_Result struct {
	// Value returned by echoI16 after a successful execution.
	Success *int16 `json:"success,omitempty"`
}

// ToWire translates a Echo_EchoI16_Result struct into a Thrift-level intermediate
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
func (v *Echo_EchoI16_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueI16(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoI16_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_EchoI16_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_EchoI16_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_EchoI16_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_EchoI16_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI16 {
				var x int16
				x, err = field.Value.GetI16(), error(nil)
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
		return fmt.Errorf("Echo_EchoI16_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Echo_EchoI16_Result
// struct.
func (v *Echo_EchoI16_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("Echo_EchoI16_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_EchoI16_Result match the
// provided Echo_EchoI16_Result.
//
// This function performs a deep comparison.
func (v *Echo_EchoI16_Result) Equals(rhs *Echo_EchoI16_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_I16_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_EchoI16_Result.
func (v *Echo_EchoI16_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddInt16("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Echo_EchoI16_Result) GetSuccess() (o int16) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Echo_EchoI16_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoI16" for this struct.
func (v *Echo_EchoI16_Result) MethodName() string {
	return "echoI16"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Echo_EchoI16_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
