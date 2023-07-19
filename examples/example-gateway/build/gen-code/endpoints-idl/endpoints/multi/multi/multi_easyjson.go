// Code generated by zanzibar
// @generated
// Checksum : G6zr21893lKn4OPpl7Ua3g==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package multi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(in *jlexer.Lexer, out *ServiceCFront_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(out *jwriter.Writer, in ServiceCFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(in *jlexer.Lexer, out *ServiceCFront_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(out *jwriter.Writer, in ServiceCFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceCFrontHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(in *jlexer.Lexer, out *ServiceBFront_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(out *jwriter.Writer, in ServiceBFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(in *jlexer.Lexer, out *ServiceBFront_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(out *jwriter.Writer, in ServiceBFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceBFrontHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(in *jlexer.Lexer, out *ServiceAFront_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(out *jwriter.Writer, in ServiceAFront_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceAFront_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceAFront_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceAFront_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceAFront_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(in *jlexer.Lexer, out *ServiceAFront_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(out *jwriter.Writer, in ServiceAFront_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceAFront_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceAFront_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceAFront_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceAFront_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarV2ExamplesExampleGatewayBuildGenCodeEndpointsIdlEndpointsMultiMultiServiceAFrontHello1(l, v)
}
