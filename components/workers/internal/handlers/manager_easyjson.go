// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

import (
	json "encoding/json"
	time "time"

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

func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(in *jlexer.Lexer, out *managerHandlerResponseSuccess) {
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
		case "result":
			out.Result = string(in.String())
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(out *jwriter.Writer, in managerHandlerResponseSuccess) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix[1:])
		out.String(string(in.Result))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(in *jlexer.Lexer, out *managerHandlerItemWorker) {
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
		case "id":
			out.ID = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
		case "locked":
			out.Locked = bool(in.Bool())
		case "task":
			if in.IsNull() {
				in.Skip()
				out.Task = nil
			} else {
				if out.Task == nil {
					out.Task = new(managerHandlerItemTask)
				}
				(*out.Task).UnmarshalEasyJSON(in)
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(out *jwriter.Writer, in managerHandlerItemWorker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"locked\":"
		out.RawString(prefix)
		out.Bool(bool(in.Locked))
	}
	{
		const prefix string = ",\"task\":"
		out.RawString(prefix)
		if in.Task == nil {
			out.RawString("null")
		} else {
			(*in.Task).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemWorker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemWorker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(in *jlexer.Lexer, out *managerHandlerItemTask) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "priority":
			out.Priority = int64(in.Int64())
		case "repeats":
			out.Repeats = int64(in.Int64())
		case "repeat_interval":
			out.RepeatInterval = time.Duration(in.Int64())
		case "timeout":
			out.Timeout = time.Duration(in.Int64())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "started_at":
			if in.IsNull() {
				in.Skip()
				out.StartedAt = nil
			} else {
				if out.StartedAt == nil {
					out.StartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.StartedAt).UnmarshalJSON(data))
				}
			}
		case "status":
			out.Status = string(in.String())
		case "locked":
			out.Locked = bool(in.Bool())
		case "attempts":
			out.Attempts = int64(in.Int64())
		case "allow_start_at":
			if in.IsNull() {
				in.Skip()
				out.AllowStartAt = nil
			} else {
				if out.AllowStartAt == nil {
					out.AllowStartAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.AllowStartAt).UnmarshalJSON(data))
				}
			}
		case "first_started_at":
			if in.IsNull() {
				in.Skip()
				out.FirstStartedAt = nil
			} else {
				if out.FirstStartedAt == nil {
					out.FirstStartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.FirstStartedAt).UnmarshalJSON(data))
				}
			}
		case "last_started_at":
			if in.IsNull() {
				in.Skip()
				out.LastStartedAt = nil
			} else {
				if out.LastStartedAt == nil {
					out.LastStartedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastStartedAt).UnmarshalJSON(data))
				}
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(out *jwriter.Writer, in managerHandlerItemTask) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"priority\":"
		out.RawString(prefix)
		out.Int64(int64(in.Priority))
	}
	{
		const prefix string = ",\"repeats\":"
		out.RawString(prefix)
		out.Int64(int64(in.Repeats))
	}
	{
		const prefix string = ",\"repeat_interval\":"
		out.RawString(prefix)
		out.Int64(int64(in.RepeatInterval))
	}
	{
		const prefix string = ",\"timeout\":"
		out.RawString(prefix)
		out.Int64(int64(in.Timeout))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"started_at\":"
		out.RawString(prefix)
		if in.StartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.StartedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"locked\":"
		out.RawString(prefix)
		out.Bool(bool(in.Locked))
	}
	{
		const prefix string = ",\"attempts\":"
		out.RawString(prefix)
		out.Int64(int64(in.Attempts))
	}
	{
		const prefix string = ",\"allow_start_at\":"
		out.RawString(prefix)
		if in.AllowStartAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.AllowStartAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"first_started_at\":"
		out.RawString(prefix)
		if in.FirstStartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.FirstStartedAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"last_started_at\":"
		out.RawString(prefix)
		if in.LastStartedAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastStartedAt).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemTask) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemTask) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(in *jlexer.Lexer, out *managerHandlerItemListener) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "locked":
			out.Locked = bool(in.Bool())
		case "events":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Events = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Events)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "fires":
			out.Fires = int64(in.Int64())
		case "first_fired_at":
			if in.IsNull() {
				in.Skip()
				out.FirstFiredAt = nil
			} else {
				if out.FirstFiredAt == nil {
					out.FirstFiredAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.FirstFiredAt).UnmarshalJSON(data))
				}
			}
		case "last_fired_at":
			if in.IsNull() {
				in.Skip()
				out.LastFiredAt = nil
			} else {
				if out.LastFiredAt == nil {
					out.LastFiredAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastFiredAt).UnmarshalJSON(data))
				}
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(out *jwriter.Writer, in managerHandlerItemListener) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"locked\":"
		out.RawString(prefix)
		out.Bool(bool(in.Locked))
	}
	{
		const prefix string = ",\"events\":"
		out.RawString(prefix)
		if in.Events == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Events {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"fires\":"
		out.RawString(prefix)
		out.Int64(int64(in.Fires))
	}
	{
		const prefix string = ",\"first_fired_at\":"
		out.RawString(prefix)
		if in.FirstFiredAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.FirstFiredAt).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"last_fired_at\":"
		out.RawString(prefix)
		if in.LastFiredAt == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.LastFiredAt).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemListener) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemListener) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(l, v)
}
