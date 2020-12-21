// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson2f147938DecodeGithubComGoParkMailRu20202CodeExpressInternalModels(in *jlexer.Lexer, out *Search) {
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
		case "albums":
			if in.IsNull() {
				in.Skip()
				out.Albums = nil
			} else {
				in.Delim('[')
				if out.Albums == nil {
					if !in.IsDelim(']') {
						out.Albums = make([]*Album, 0, 8)
					} else {
						out.Albums = []*Album{}
					}
				} else {
					out.Albums = (out.Albums)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Album
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Album)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Albums = append(out.Albums, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "artists":
			if in.IsNull() {
				in.Skip()
				out.Artists = nil
			} else {
				in.Delim('[')
				if out.Artists == nil {
					if !in.IsDelim(']') {
						out.Artists = make([]*Artist, 0, 8)
					} else {
						out.Artists = []*Artist{}
					}
				} else {
					out.Artists = (out.Artists)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *Artist
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(Artist)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Artists = append(out.Artists, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tracks":
			if in.IsNull() {
				in.Skip()
				out.Tracks = nil
			} else {
				in.Delim('[')
				if out.Tracks == nil {
					if !in.IsDelim(']') {
						out.Tracks = make([]*Track, 0, 8)
					} else {
						out.Tracks = []*Track{}
					}
				} else {
					out.Tracks = (out.Tracks)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *Track
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(Track)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.Tracks = append(out.Tracks, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]*User, 0, 8)
					} else {
						out.Users = []*User{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *User
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(User)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Users = append(out.Users, v4)
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
func easyjson2f147938EncodeGithubComGoParkMailRu20202CodeExpressInternalModels(out *jwriter.Writer, in Search) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"albums\":"
		out.RawString(prefix[1:])
		if in.Albums == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Albums {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"artists\":"
		out.RawString(prefix)
		if in.Artists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Artists {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					(*v8).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tracks\":"
		out.RawString(prefix)
		if in.Tracks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Tracks {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Users {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					(*v12).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Search) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2f147938EncodeGithubComGoParkMailRu20202CodeExpressInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Search) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2f147938EncodeGithubComGoParkMailRu20202CodeExpressInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Search) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2f147938DecodeGithubComGoParkMailRu20202CodeExpressInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Search) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2f147938DecodeGithubComGoParkMailRu20202CodeExpressInternalModels(l, v)
}
