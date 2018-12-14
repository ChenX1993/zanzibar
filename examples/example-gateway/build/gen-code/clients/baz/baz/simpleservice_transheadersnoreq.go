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

// SimpleService_TransHeadersNoReq_Args represents the arguments for the SimpleService.transHeadersNoReq function.
//
// The arguments for transHeadersNoReq are sent and received over the wire as this struct.
type SimpleService_TransHeadersNoReq_Args struct {
	Req *base.NestedStruct `json:"req,required"`
	S2  *string            `json:"s2,omitempty"`
	I3  int32              `json:"i3,required"`
	B4  *bool              `json:"b4,omitempty"`
}

// ToWire translates a SimpleService_TransHeadersNoReq_Args struct into a Thrift-level intermediate
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
func (v *SimpleService_TransHeadersNoReq_Args) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Req == nil {
		return w, errors.New("field Req of SimpleService_TransHeadersNoReq_Args is required")
	}
	w, err = v.Req.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.S2 != nil {
		w, err = wire.NewValueString(*(v.S2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	w, err = wire.NewValueI32(v.I3), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.B4 != nil {
		w, err = wire.NewValueBool(*(v.B4)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _NestedStruct_Read(w wire.Value) (*base.NestedStruct, error) {
	var v base.NestedStruct
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a SimpleService_TransHeadersNoReq_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_TransHeadersNoReq_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_TransHeadersNoReq_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_TransHeadersNoReq_Args) FromWire(w wire.Value) error {
	var err error

	reqIsSet := false

	i3IsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Req, err = _NestedStruct_Read(field.Value)
				if err != nil {
					return err
				}
				reqIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.S2 = &x
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TI32 {
				v.I3, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				i3IsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.B4 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if !reqIsSet {
		return errors.New("field Req of SimpleService_TransHeadersNoReq_Args is required")
	}

	if !i3IsSet {
		return errors.New("field I3 of SimpleService_TransHeadersNoReq_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SimpleService_TransHeadersNoReq_Args
// struct.
func (v *SimpleService_TransHeadersNoReq_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [4]string
	i := 0
	fields[i] = fmt.Sprintf("Req: %v", v.Req)
	i++
	if v.S2 != nil {
		fields[i] = fmt.Sprintf("S2: %v", *(v.S2))
		i++
	}
	fields[i] = fmt.Sprintf("I3: %v", v.I3)
	i++
	if v.B4 != nil {
		fields[i] = fmt.Sprintf("B4: %v", *(v.B4))
		i++
	}

	return fmt.Sprintf("SimpleService_TransHeadersNoReq_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_TransHeadersNoReq_Args match the
// provided SimpleService_TransHeadersNoReq_Args.
//
// This function performs a deep comparison.
func (v *SimpleService_TransHeadersNoReq_Args) Equals(rhs *SimpleService_TransHeadersNoReq_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Req.Equals(rhs.Req) {
		return false
	}
	if !_String_EqualsPtr(v.S2, rhs.S2) {
		return false
	}
	if !(v.I3 == rhs.I3) {
		return false
	}
	if !_Bool_EqualsPtr(v.B4, rhs.B4) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_TransHeadersNoReq_Args.
func (v *SimpleService_TransHeadersNoReq_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("req", v.Req))
	if v.S2 != nil {
		enc.AddString("s2", *v.S2)
	}
	enc.AddInt32("i3", v.I3)
	if v.B4 != nil {
		enc.AddBool("b4", *v.B4)
	}
	return err
}

// GetReq returns the value of Req if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Args) GetReq() (o *base.NestedStruct) { return v.Req }

// IsSetReq returns true if Req is not nil.
func (v *SimpleService_TransHeadersNoReq_Args) IsSetReq() bool {
	return v.Req != nil
}

// GetS2 returns the value of S2 if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Args) GetS2() (o string) {
	if v.S2 != nil {
		return *v.S2
	}

	return
}

// IsSetS2 returns true if S2 is not nil.
func (v *SimpleService_TransHeadersNoReq_Args) IsSetS2() bool {
	return v.S2 != nil
}

// GetI3 returns the value of I3 if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Args) GetI3() (o int32) { return v.I3 }

// GetB4 returns the value of B4 if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Args) GetB4() (o bool) {
	if v.B4 != nil {
		return *v.B4
	}

	return
}

// IsSetB4 returns true if B4 is not nil.
func (v *SimpleService_TransHeadersNoReq_Args) IsSetB4() bool {
	return v.B4 != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "transHeadersNoReq" for this struct.
func (v *SimpleService_TransHeadersNoReq_Args) MethodName() string {
	return "transHeadersNoReq"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SimpleService_TransHeadersNoReq_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SimpleService_TransHeadersNoReq_Helper provides functions that aid in handling the
// parameters and return values of the SimpleService.transHeadersNoReq
// function.
var SimpleService_TransHeadersNoReq_Helper = struct {
	// Args accepts the parameters of transHeadersNoReq in-order and returns
	// the arguments struct for the function.
	Args func(
		req *base.NestedStruct,
		s2 *string,
		i3 int32,
		b4 *bool,
	) *SimpleService_TransHeadersNoReq_Args

	// IsException returns true if the given error can be thrown
	// by transHeadersNoReq.
	//
	// An error can be thrown by transHeadersNoReq only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for transHeadersNoReq
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// transHeadersNoReq into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by transHeadersNoReq
	//
	//   value, err := transHeadersNoReq(args)
	//   result, err := SimpleService_TransHeadersNoReq_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from transHeadersNoReq: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*base.TransHeaders, error) (*SimpleService_TransHeadersNoReq_Result, error)

	// UnwrapResponse takes the result struct for transHeadersNoReq
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if transHeadersNoReq threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SimpleService_TransHeadersNoReq_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SimpleService_TransHeadersNoReq_Result) (*base.TransHeaders, error)
}{}

func init() {
	SimpleService_TransHeadersNoReq_Helper.Args = func(
		req *base.NestedStruct,
		s2 *string,
		i3 int32,
		b4 *bool,
	) *SimpleService_TransHeadersNoReq_Args {
		return &SimpleService_TransHeadersNoReq_Args{
			Req: req,
			S2:  s2,
			I3:  i3,
			B4:  b4,
		}
	}

	SimpleService_TransHeadersNoReq_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		default:
			return false
		}
	}

	SimpleService_TransHeadersNoReq_Helper.WrapResponse = func(success *base.TransHeaders, err error) (*SimpleService_TransHeadersNoReq_Result, error) {
		if err == nil {
			return &SimpleService_TransHeadersNoReq_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_TransHeadersNoReq_Result.AuthErr")
			}
			return &SimpleService_TransHeadersNoReq_Result{AuthErr: e}, nil
		}

		return nil, err
	}
	SimpleService_TransHeadersNoReq_Helper.UnwrapResponse = func(result *SimpleService_TransHeadersNoReq_Result) (success *base.TransHeaders, err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
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

// SimpleService_TransHeadersNoReq_Result represents the result of a SimpleService.transHeadersNoReq function call.
//
// The result of a transHeadersNoReq execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SimpleService_TransHeadersNoReq_Result struct {
	// Value returned by transHeadersNoReq after a successful execution.
	Success *base.TransHeaders `json:"success,omitempty"`
	AuthErr *AuthErr           `json:"authErr,omitempty"`
}

// ToWire translates a SimpleService_TransHeadersNoReq_Result struct into a Thrift-level intermediate
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
func (v *SimpleService_TransHeadersNoReq_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
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

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_TransHeadersNoReq_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SimpleService_TransHeadersNoReq_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SimpleService_TransHeadersNoReq_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SimpleService_TransHeadersNoReq_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SimpleService_TransHeadersNoReq_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _TransHeaders_Read(field.Value)
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
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.AuthErr != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_TransHeadersNoReq_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SimpleService_TransHeadersNoReq_Result
// struct.
func (v *SimpleService_TransHeadersNoReq_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}

	return fmt.Sprintf("SimpleService_TransHeadersNoReq_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SimpleService_TransHeadersNoReq_Result match the
// provided SimpleService_TransHeadersNoReq_Result.
//
// This function performs a deep comparison.
func (v *SimpleService_TransHeadersNoReq_Result) Equals(rhs *SimpleService_TransHeadersNoReq_Result) bool {
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

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SimpleService_TransHeadersNoReq_Result.
func (v *SimpleService_TransHeadersNoReq_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.AuthErr != nil {
		err = multierr.Append(err, enc.AddObject("authErr", v.AuthErr))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Result) GetSuccess() (o *base.TransHeaders) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *SimpleService_TransHeadersNoReq_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// GetAuthErr returns the value of AuthErr if it is set or its
// zero value if it is unset.
func (v *SimpleService_TransHeadersNoReq_Result) GetAuthErr() (o *AuthErr) {
	if v.AuthErr != nil {
		return v.AuthErr
	}

	return
}

// IsSetAuthErr returns true if AuthErr is not nil.
func (v *SimpleService_TransHeadersNoReq_Result) IsSetAuthErr() bool {
	return v.AuthErr != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "transHeadersNoReq" for this struct.
func (v *SimpleService_TransHeadersNoReq_Result) MethodName() string {
	return "transHeadersNoReq"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SimpleService_TransHeadersNoReq_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
