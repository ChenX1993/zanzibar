// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// SimpleService_Trans_Args represents the arguments for the SimpleService.trans function.
//
// The arguments for trans are sent and received over the wire as this struct.
type SimpleService_Trans_Args struct {
	Arg1 *base.TransStruct `json:"arg1,required"`
	Arg2 *base.TransStruct `json:"arg2,omitempty"`
}

// ToWire translates a SimpleService_Trans_Args struct into a Thrift-level intermediate
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
func (v *SimpleService_Trans_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg1 == nil {
		return w, errors.New("field Arg1 of SimpleService_Trans_Args is required")
	}
	w, err = v.Arg1.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Arg2 != nil {
		w, err = v.Arg2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _TransStruct_Read(w wire.Value) (*base.TransStruct, error) {
	var v base.TransStruct
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a SimpleService_Trans_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Trans_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Trans_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Trans_Args) FromWire(w wire.Value) error {
	var err error

	arg1IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Arg1, err = _TransStruct_Read(field.Value)
				if err != nil {
					return err
				}
				arg1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Arg2, err = _TransStruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !arg1IsSet {
		return errors.New("field Arg1 of SimpleService_Trans_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Trans_Args
// struct.
func (v *SimpleService_Trans_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Arg1: %v", v.Arg1)
	i++
	if v.Arg2 != nil {
		fields[i] = fmt.Sprintf("Arg2: %v", v.Arg2)
		i++
	}

	return fmt.Sprintf("SimpleService_Trans_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Trans_Args match the
// provided SimpleService_Trans_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_Trans_Args) Equals(rhs *SimpleService_Trans_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Arg1.Equals(rhs.Arg1) {
		return false
	}
	if !((v.Arg2 == nil && rhs.Arg2 == nil) || (v.Arg2 != nil && rhs.Arg2 != nil && v.Arg2.Equals(rhs.Arg2))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_Trans_Args.
func (v *SimpleService_Trans_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("arg1", v.Arg1))
	if v.Arg2 != nil {
		err = multierr.Append(err, enc.AddObject("arg2", v.Arg2))
	}
	return err
}

// GetArg1 returns the value of Arg1 if it is set or its
// zero value if it is unset.
func (v *SimpleService_Trans_Args) GetArg1() (o *base.TransStruct) { return v.Arg1 }

// IsSetArg1 returns true if Arg1 is not nil.
func (v *SimpleService_Trans_Args) IsSetArg1() bool {
	return v.Arg1 != nil
}

// GetArg2 returns the value of Arg2 if it is set or its
// zero value if it is unset.
func (v *SimpleService_Trans_Args) GetArg2() (o *base.TransStruct) {
	if v.Arg2 != nil {
		return v.Arg2
	}

	return
}

// IsSetArg2 returns true if Arg2 is not nil.
func (v *SimpleService_Trans_Args) IsSetArg2() bool {
	return v.Arg2 != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "trans" for this struct.
func (v *SimpleService_Trans_Args) MethodName() string {
	return "trans"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_Trans_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_Trans_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.trans
// function.
var SimpleService_Trans_Helper = struct {
	// Args accepts the parameters of trans in-order and returns
	// the arguments struct for the function.
	Args func(
		arg1 *base.TransStruct,
		arg2 *base.TransStruct,
	) *SimpleService_Trans_Args

	// IsException returns true if the given error can be thrown
	// by trans.
	//
	// An error can be thrown by trans only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for trans
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// trans into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by trans
	//
	//   value, err := trans(args)
	//   result, err := SimpleService_Trans_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from trans: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*base.TransStruct, error) (*SimpleService_Trans_Result, error)

	// UnwrapResponse takes the result struct for trans
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if trans threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SimpleService_Trans_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_Trans_Result) (*base.TransStruct, error)
}{}

func init() {
	SimpleService_Trans_Helper.Args = func(
		arg1 *base.TransStruct,
		arg2 *base.TransStruct,
	) *SimpleService_Trans_Args {
		return &SimpleService_Trans_Args{
			Arg1: arg1,
			Arg2: arg2,
		}
	}

	SimpleService_Trans_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *OtherAuthErr:
			return true
		default:
			return false
		}
	}

	SimpleService_Trans_Helper.WrapResponse = func(success *base.TransStruct, err error) (*SimpleService_Trans_Result, error) {
		if err == nil {
			return &SimpleService_Trans_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Trans_Result.AuthErr")
			}
			return &SimpleService_Trans_Result{AuthErr: e}, nil
		case *OtherAuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Trans_Result.OtherAuthErr")
			}
			return &SimpleService_Trans_Result{OtherAuthErr: e}, nil
		}

		return nil, err
	}
	SimpleService_Trans_Helper.UnwrapResponse = func(result *SimpleService_Trans_Result) (success *base.TransStruct, err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.OtherAuthErr != nil {
			err = result.OtherAuthErr
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SimpleService_Trans_Result represents the result of a SimpleService.trans function call.
//
// The result of a trans execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SimpleService_Trans_Result struct {
	// Value returned by trans after a successful execution.
	Success      *base.TransStruct `json:"success,omitempty"`
	AuthErr      *AuthErr          `json:"authErr,omitempty"`
	OtherAuthErr *OtherAuthErr     `json:"otherAuthErr,omitempty"`
}

// ToWire translates a SimpleService_Trans_Result struct into a Thrift-level intermediate
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
func (v *SimpleService_Trans_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
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
	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.OtherAuthErr != nil {
		w, err = v.OtherAuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_Trans_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_Trans_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_Trans_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_Trans_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_Trans_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _TransStruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.OtherAuthErr, err = _OtherAuthErr_Read(field.Value)
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
	if v.AuthErr != nil {
		count++
	}
	if v.OtherAuthErr != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_Trans_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_Trans_Result
// struct.
func (v *SimpleService_Trans_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.OtherAuthErr != nil {
		fields[i] = fmt.Sprintf("OtherAuthErr: %v", v.OtherAuthErr)
		i++
	}

	return fmt.Sprintf("SimpleService_Trans_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_Trans_Result match the
// provided SimpleService_Trans_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_Trans_Result) Equals(rhs *SimpleService_Trans_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.OtherAuthErr == nil && rhs.OtherAuthErr == nil) || (v.OtherAuthErr != nil && rhs.OtherAuthErr != nil && v.OtherAuthErr.Equals(rhs.OtherAuthErr))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_Trans_Result.
func (v *SimpleService_Trans_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.AuthErr != nil {
		err = multierr.Append(err, enc.AddObject("authErr", v.AuthErr))
	}
	if v.OtherAuthErr != nil {
		err = multierr.Append(err, enc.AddObject("otherAuthErr", v.OtherAuthErr))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SimpleService_Trans_Result) GetSuccess() (o *base.TransStruct) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *SimpleService_Trans_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// GetAuthErr returns the value of AuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_Trans_Result) GetAuthErr() (o *AuthErr) {
	if v.AuthErr != nil {
		return v.AuthErr
	}

	return
}

// IsSetAuthErr returns true if AuthErr is not nil.
func (v *SimpleService_Trans_Result) IsSetAuthErr() bool {
	return v.AuthErr != nil
}

// GetOtherAuthErr returns the value of OtherAuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_Trans_Result) GetOtherAuthErr() (o *OtherAuthErr) {
	if v.OtherAuthErr != nil {
		return v.OtherAuthErr
	}

	return
}

// IsSetOtherAuthErr returns true if OtherAuthErr is not nil.
func (v *SimpleService_Trans_Result) IsSetOtherAuthErr() bool {
	return v.OtherAuthErr != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "trans" for this struct.
func (v *SimpleService_Trans_Result) MethodName() string {
	return "trans"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_Trans_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
