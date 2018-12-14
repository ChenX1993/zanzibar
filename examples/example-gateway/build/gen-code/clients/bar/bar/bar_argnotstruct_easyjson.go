// Code generated by zanzibar
// @generated
// Checksum : /D/faqhaXMij6gkJeMi6gg==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(in *jlexer.Lexer, out *Bar_ArgNotStruct_Result) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "barException":
			if in.IsNull() {
				in.Skip()
				out.BarException = nil
			} else {
				if out.BarException == nil {
					out.BarException = new(BarException)
				}
				easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in, &*out.BarException)
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
func easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(out *jwriter.Writer, in Bar_ArgNotStruct_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.BarException != nil {
		const prefix string = ",\"barException\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out, *in.BarException)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgNotStruct_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgNotStruct_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgNotStruct_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgNotStruct_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct(l, v)
}
func easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(in *jlexer.Lexer, out *BarException) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
}
func easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBar(out *jwriter.Writer, in BarException) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"stringField\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StringField))
	}
	out.RawByte('}')
}
func easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(in *jlexer.Lexer, out *Bar_ArgNotStruct_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var RequestSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "request":
			out.Request = string(in.String())
			RequestSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !RequestSet {
		in.AddError(fmt.Errorf("key 'request' is required"))
	}
}
func easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(out *jwriter.Writer, in Bar_ArgNotStruct_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"request\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Request))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgNotStruct_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgNotStruct_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82fab59aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgNotStruct_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgNotStruct_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82fab59aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarBarArgNotStruct1(l, v)
}
