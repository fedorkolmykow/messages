// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package modeles

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

func easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles(in *jlexer.Lexer, out *ChatsGetResponse) {
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
		case "chats":
			if in.IsNull() {
				in.Skip()
				out.Chats = nil
			} else {
				in.Delim('[')
				if out.Chats == nil {
					if !in.IsDelim(']') {
						out.Chats = make([]Chat, 0, 0)
					} else {
						out.Chats = []Chat{}
					}
				} else {
					out.Chats = (out.Chats)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Chat
					(v1).UnmarshalEasyJSON(in)
					out.Chats = append(out.Chats, v1)
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
func easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles(out *jwriter.Writer, in ChatsGetResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chats\":"
		out.RawString(prefix[1:])
		if in.Chats == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Chats {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatsGetResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatsGetResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatsGetResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatsGetResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles(l, v)
}
func easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles1(in *jlexer.Lexer, out *ChatsGetRequest) {
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
		case "user":
			out.UserId = string(in.String())
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
func easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles1(out *jwriter.Writer, in ChatsGetRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatsGetRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatsGetRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatsGetRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatsGetRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles1(l, v)
}
func easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles2(in *jlexer.Lexer, out *ChatAddResponse) {
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
		case "id":
			out.ChatId = string(in.String())
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
func easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles2(out *jwriter.Writer, in ChatAddResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ChatId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatAddResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatAddResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatAddResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatAddResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles2(l, v)
}
func easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles3(in *jlexer.Lexer, out *ChatAddRequest) {
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
		case "name":
			out.Name = string(in.String())
		case "users":
			if in.IsNull() {
				in.Skip()
				out.UsersId = nil
			} else {
				in.Delim('[')
				if out.UsersId == nil {
					if !in.IsDelim(']') {
						out.UsersId = make([]string, 0, 4)
					} else {
						out.UsersId = []string{}
					}
				} else {
					out.UsersId = (out.UsersId)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.UsersId = append(out.UsersId, v4)
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
func easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles3(out *jwriter.Writer, in ChatAddRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.UsersId == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.UsersId {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChatAddRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChatAddRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChatAddRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChatAddRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles3(l, v)
}
func easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles4(in *jlexer.Lexer, out *Chat) {
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
		case "id":
			out.ChatId = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]string, 0, 4)
					} else {
						out.Users = []string{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Users = append(out.Users, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "created_at":
			out.CreatedAt = string(in.String())
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
func easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles4(out *jwriter.Writer, in Chat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ChatId))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Users {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.String(string(in.CreatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Chat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Chat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF77207f2EncodeAvitoMessageMessageServicePkgModeles4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Chat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Chat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF77207f2DecodeAvitoMessageMessageServicePkgModeles4(l, v)
}
