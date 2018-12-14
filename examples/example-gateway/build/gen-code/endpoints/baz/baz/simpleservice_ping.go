// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// SimpleService_Ping_Args represents the arguments for the SimpleService.ping function.
//
// The arguments for ping are sent and received over the wire as this struct.
type SimpleService_Ping_Args struct {
}

// ToWire translates a SimpleService_Ping_Args struct into a Thrift-level intermediate
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
func (v *SimpleService_Ping_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_Ping_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Ping_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Ping_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Ping_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Ping_Args
// struct.
func (v *SimpleService_Ping_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("SimpleService_Ping_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Ping_Args match the
// provided SimpleService_Ping_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_Ping_Args) Equals(rhs *SimpleService_Ping_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_Ping_Args.
func (v *SimpleService_Ping_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "ping" for this struct.
func (v *SimpleService_Ping_Args) MethodName() string {
	return "ping"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_Ping_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_Ping_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.ping
// function.
var SimpleService_Ping_Helper = struct {
	// Args accepts the parameters of ping in-order and returns
	// the arguments struct for the function.
	Args func() *SimpleService_Ping_Args

	// IsException returns true if the given error can be thrown
	// by ping.
	//
	// An error can be thrown by ping only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for ping
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// ping into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by ping
	//
	//   value, err := ping(args)
	//   result, err := SimpleService_Ping_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from ping: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BazResponse, error) (*SimpleService_Ping_Result, error)

	// UnwrapResponse takes the result struct for ping
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if ping threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SimpleService_Ping_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_Ping_Result) (*BazResponse, error)
}{}

func init() {
	SimpleService_Ping_Helper.Args = func() *SimpleService_Ping_Args {
		return &SimpleService_Ping_Args{}
	}

	SimpleService_Ping_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	SimpleService_Ping_Helper.WrapResponse = func(success *BazResponse, err error) (*SimpleService_Ping_Result, error) {
		if err == nil {
			return &SimpleService_Ping_Result{Success: success}, nil
		}

		return nil, err
	}
	SimpleService_Ping_Helper.UnwrapResponse = func(result *SimpleService_Ping_Result) (success *BazResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SimpleService_Ping_Result represents the result of a SimpleService.ping function call.
//
// The result of a ping execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SimpleService_Ping_Result struct {
	// Value returned by ping after a successful execution.
	Success *BazResponse `json:"success,omitempty"`
}

// ToWire translates a SimpleService_Ping_Result struct into a Thrift-level intermediate
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
func (v *SimpleService_Ping_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("SimpleService_Ping_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_Ping_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Ping_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Ping_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Ping_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BazResponse_Read(field.Value)
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
		return fmt.Errorf("SimpleService_Ping_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Ping_Result
// struct.
func (v *SimpleService_Ping_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("SimpleService_Ping_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Ping_Result match the
// provided SimpleService_Ping_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_Ping_Result) Equals(rhs *SimpleService_Ping_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_Ping_Result.
func (v *SimpleService_Ping_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SimpleService_Ping_Result) GetSuccess() (o *BazResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *SimpleService_Ping_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "ping" for this struct.
func (v *SimpleService_Ping_Result) MethodName() string {
	return "ping"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_Ping_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
