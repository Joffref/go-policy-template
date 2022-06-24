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

func easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *NodeAffinity) {
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
		case "preferredDuringSchedulingIgnoredDuringExecution":
			if in.IsNull() {
				in.Skip()
				out.PreferredDuringSchedulingIgnoredDuringExecution = nil
			} else {
				in.Delim('[')
				if out.PreferredDuringSchedulingIgnoredDuringExecution == nil {
					if !in.IsDelim(']') {
						out.PreferredDuringSchedulingIgnoredDuringExecution = make([]*PreferredSchedulingTerm, 0, 8)
					} else {
						out.PreferredDuringSchedulingIgnoredDuringExecution = []*PreferredSchedulingTerm{}
					}
				} else {
					out.PreferredDuringSchedulingIgnoredDuringExecution = (out.PreferredDuringSchedulingIgnoredDuringExecution)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PreferredSchedulingTerm
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PreferredSchedulingTerm)
						}
						easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, v1)
					}
					out.PreferredDuringSchedulingIgnoredDuringExecution = append(out.PreferredDuringSchedulingIgnoredDuringExecution, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "requiredDuringSchedulingIgnoredDuringExecution":
			if in.IsNull() {
				in.Skip()
				out.RequiredDuringSchedulingIgnoredDuringExecution = nil
			} else {
				if out.RequiredDuringSchedulingIgnoredDuringExecution == nil {
					out.RequiredDuringSchedulingIgnoredDuringExecution = new(NodeSelector)
				}
				easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.RequiredDuringSchedulingIgnoredDuringExecution)
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
func easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in NodeAffinity) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"preferredDuringSchedulingIgnoredDuringExecution\":"
		out.RawString(prefix[1:])
		if in.PreferredDuringSchedulingIgnoredDuringExecution == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.PreferredDuringSchedulingIgnoredDuringExecution {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		const prefix string = ",\"requiredDuringSchedulingIgnoredDuringExecution\":"
		out.RawString(prefix)
		easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.RequiredDuringSchedulingIgnoredDuringExecution)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeAffinity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeAffinity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeAffinity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeAffinity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *NodeSelector) {
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
		case "nodeSelectorTerms":
			if in.IsNull() {
				in.Skip()
				out.NodeSelectorTerms = nil
			} else {
				in.Delim('[')
				if out.NodeSelectorTerms == nil {
					if !in.IsDelim(']') {
						out.NodeSelectorTerms = make([]*NodeSelectorTerm, 0, 8)
					} else {
						out.NodeSelectorTerms = []*NodeSelectorTerm{}
					}
				} else {
					out.NodeSelectorTerms = (out.NodeSelectorTerms)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *NodeSelectorTerm
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(NodeSelectorTerm)
						}
						easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, v4)
					}
					out.NodeSelectorTerms = append(out.NodeSelectorTerms, v4)
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
func easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in NodeSelector) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nodeSelectorTerms\":"
		out.RawString(prefix[1:])
		if in.NodeSelectorTerms == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.NodeSelectorTerms {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *NodeSelectorTerm) {
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
		case "matchExpressions":
			if in.IsNull() {
				in.Skip()
				out.MatchExpressions = nil
			} else {
				in.Delim('[')
				if out.MatchExpressions == nil {
					if !in.IsDelim(']') {
						out.MatchExpressions = make([]*NodeSelectorRequirement, 0, 8)
					} else {
						out.MatchExpressions = []*NodeSelectorRequirement{}
					}
				} else {
					out.MatchExpressions = (out.MatchExpressions)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *NodeSelectorRequirement
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(NodeSelectorRequirement)
						}
						easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, v7)
					}
					out.MatchExpressions = append(out.MatchExpressions, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "matchFields":
			if in.IsNull() {
				in.Skip()
				out.MatchFields = nil
			} else {
				in.Delim('[')
				if out.MatchFields == nil {
					if !in.IsDelim(']') {
						out.MatchFields = make([]*NodeSelectorRequirement, 0, 8)
					} else {
						out.MatchFields = []*NodeSelectorRequirement{}
					}
				} else {
					out.MatchFields = (out.MatchFields)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *NodeSelectorRequirement
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(NodeSelectorRequirement)
						}
						easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, v8)
					}
					out.MatchFields = append(out.MatchFields, v8)
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
func easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in NodeSelectorTerm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"matchExpressions\":"
		out.RawString(prefix[1:])
		if in.MatchExpressions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.MatchExpressions {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *v10)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"matchFields\":"
		out.RawString(prefix)
		if in.MatchFields == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.MatchFields {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *v12)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV14(in *jlexer.Lexer, out *NodeSelectorRequirement) {
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
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(string)
				}
				*out.Key = string(in.String())
			}
		case "operator":
			if in.IsNull() {
				in.Skip()
				out.Operator = nil
			} else {
				if out.Operator == nil {
					out.Operator = new(string)
				}
				*out.Operator = string(in.String())
			}
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]string, 0, 4)
					} else {
						out.Values = []string{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Values = append(out.Values, v13)
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
func easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV14(out *jwriter.Writer, in NodeSelectorRequirement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		if in.Key == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Key))
		}
	}
	{
		const prefix string = ",\"operator\":"
		out.RawString(prefix)
		if in.Operator == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Operator))
		}
	}
	{
		const prefix string = ",\"values\":"
		out.RawString(prefix)
		if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Values {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *PreferredSchedulingTerm) {
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
		case "preference":
			if in.IsNull() {
				in.Skip()
				out.Preference = nil
			} else {
				if out.Preference == nil {
					out.Preference = new(NodeSelectorTerm)
				}
				easyjson2275aad5DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, out.Preference)
			}
		case "weight":
			if in.IsNull() {
				in.Skip()
				out.Weight = nil
			} else {
				if out.Weight == nil {
					out.Weight = new(int32)
				}
				*out.Weight = int32(in.Int32())
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
func easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in PreferredSchedulingTerm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"preference\":"
		out.RawString(prefix[1:])
		if in.Preference == nil {
			out.RawString("null")
		} else {
			easyjson2275aad5EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *in.Preference)
		}
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		if in.Weight == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Weight))
		}
	}
	out.RawByte('}')
}