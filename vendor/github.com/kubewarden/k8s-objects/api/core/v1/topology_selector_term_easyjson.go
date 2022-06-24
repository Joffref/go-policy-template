// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonB66dba6eDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *TopologySelectorTerm) {
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
		case "matchLabelExpressions":
			if in.IsNull() {
				in.Skip()
				out.MatchLabelExpressions = nil
			} else {
				in.Delim('[')
				if out.MatchLabelExpressions == nil {
					if !in.IsDelim(']') {
						out.MatchLabelExpressions = make([]*TopologySelectorLabelRequirement, 0, 8)
					} else {
						out.MatchLabelExpressions = []*TopologySelectorLabelRequirement{}
					}
				} else {
					out.MatchLabelExpressions = (out.MatchLabelExpressions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *TopologySelectorLabelRequirement
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(TopologySelectorLabelRequirement)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.MatchLabelExpressions = append(out.MatchLabelExpressions, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonB66dba6eEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in TopologySelectorTerm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"matchLabelExpressions\":"
		out.RawString(prefix[1:])
		if in.MatchLabelExpressions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MatchLabelExpressions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopologySelectorTerm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB66dba6eEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopologySelectorTerm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB66dba6eEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopologySelectorTerm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB66dba6eDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopologySelectorTerm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB66dba6eDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}