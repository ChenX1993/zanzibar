// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package corge

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Corge_EchoBool_Args represents the arguments for the Corge.echoBool function.
//
// The arguments for echoBool are sent and received over the wire as this struct.
type Corge_EchoBool_Args struct {
	Arg bool `json:"arg,required"`
}

// ToWire translates a Corge_EchoBool_Args struct into a Thrift-level intermediate
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
func (v *Corge_EchoBool_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueBool(v.Arg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Corge_EchoBool_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Corge_EchoBool_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Corge_EchoBool_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Corge_EchoBool_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				v.Arg, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of Corge_EchoBool_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Corge_EchoBool_Args
// struct.
func (v *Corge_EchoBool_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("Corge_EchoBool_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Corge_EchoBool_Args match the
// provided Corge_EchoBool_Args.
//
// This function performs a deep comparison.
func (v *Corge_EchoBool_Args) Equals(rhs *Corge_EchoBool_Args) bool {
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
// fast logging of Corge_EchoBool_Args.
func (v *Corge_EchoBool_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddBool("arg", v.Arg)
	return err
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *Corge_EchoBool_Args) GetArg() (o bool) { return v.Arg }

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoBool" for this struct.
func (v *Corge_EchoBool_Args) MethodName() string {
	return "echoBool"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Corge_EchoBool_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Corge_EchoBool_Helper provides functions that aid in handling the
// parameters and return values of the Corge.echoBool
// function.
var Corge_EchoBool_Helper = struct {
	// Args accepts the parameters of echoBool in-order and returns
	// the arguments struct for the function.
	Args func(
		arg bool,
	) *Corge_EchoBool_Args

	// IsException returns true if the given error can be thrown
	// by echoBool.
	//
	// An error can be thrown by echoBool only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoBool
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoBool into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoBool
	//
	//   value, err := echoBool(args)
	//   result, err := Corge_EchoBool_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoBool: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(bool, error) (*Corge_EchoBool_Result, error)

	// UnwrapResponse takes the result struct for echoBool
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoBool threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Corge_EchoBool_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Corge_EchoBool_Result) (bool, error)
}{}

func init() {
	Corge_EchoBool_Helper.Args = func(
		arg bool,
	) *Corge_EchoBool_Args {
		return &Corge_EchoBool_Args{
			Arg: arg,
		}
	}

	Corge_EchoBool_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Corge_EchoBool_Helper.WrapResponse = func(success bool, err error) (*Corge_EchoBool_Result, error) {
		if err == nil {
			return &Corge_EchoBool_Result{Success: &success}, nil
		}

		return nil, err
	}
	Corge_EchoBool_Helper.UnwrapResponse = func(result *Corge_EchoBool_Result) (success bool, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Corge_EchoBool_Result represents the result of a Corge.echoBool function call.
//
// The result of a echoBool execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Corge_EchoBool_Result struct {
	// Value returned by echoBool after a successful execution.
	Success *bool `json:"success,omitempty"`
}

// ToWire translates a Corge_EchoBool_Result struct into a Thrift-level intermediate
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
func (v *Corge_EchoBool_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueBool(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Corge_EchoBool_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Corge_EchoBool_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Corge_EchoBool_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Corge_EchoBool_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Corge_EchoBool_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
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
		return fmt.Errorf("Corge_EchoBool_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Corge_EchoBool_Result
// struct.
func (v *Corge_EchoBool_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("Corge_EchoBool_Result{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Corge_EchoBool_Result match the
// provided Corge_EchoBool_Result.
//
// This function performs a deep comparison.
func (v *Corge_EchoBool_Result) Equals(rhs *Corge_EchoBool_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Bool_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Corge_EchoBool_Result.
func (v *Corge_EchoBool_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddBool("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Corge_EchoBool_Result) GetSuccess() (o bool) {
	if v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Corge_EchoBool_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoBool" for this struct.
func (v *Corge_EchoBool_Result) MethodName() string {
	return "echoBool"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Corge_EchoBool_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
