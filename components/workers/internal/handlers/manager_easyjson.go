// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

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

func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(in *jlexer.Lexer, out *managerHandlerItemListener) {
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
		case "event":
			out.Event = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(out *jwriter.Writer, in managerHandlerItemListener) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"event\":")
	out.String(string(in.Event))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemListener) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemListener) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemListener) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(in *jlexer.Lexer, out *managerHandlerItemTask) {
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
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "priority":
			out.Priority = int64(in.Int64())
		case "attempts":
			out.Attempts = int64(in.Int64())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(out *jwriter.Writer, in managerHandlerItemTask) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"status\":")
	out.String(string(in.Status))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"priority\":")
	out.Int64(int64(in.Priority))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"attempts\":")
	out.Int64(int64(in.Attempts))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"created\":")
	out.Raw((in.Created).MarshalJSON())
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemTask) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemTask) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemTask) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers1(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(in *jlexer.Lexer, out *managerHandlerItemWorker) {
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
			out.Id = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(out *jwriter.Writer, in managerHandlerItemWorker) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"created\":")
	out.Raw((in.Created).MarshalJSON())
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"status\":")
	out.String(string(in.Status))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"task\":")
	if in.Task == nil {
		out.RawString("null")
	} else {
		(*in.Task).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerItemWorker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerItemWorker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerItemWorker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers2(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(in *jlexer.Lexer, out *managerHandlerStats) {
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
		case "tasks_count":
			out.TasksCount = uint64(in.Uint64())
		case "tasks_waiting_count":
			out.TasksWaitingCount = uint64(in.Uint64())
		case "workers_count":
			out.WorkersCount = uint64(in.Uint64())
		case "workers_waiting_count":
			out.WorkersWaitingCount = uint64(in.Uint64())
		case "listeners_count":
			out.ListenersCount = uint64(in.Uint64())
		case "workers":
			if in.IsNull() {
				in.Skip()
				out.Workers = nil
			} else {
				in.Delim('[')
				if out.Workers == nil {
					if !in.IsDelim(']') {
						out.Workers = make([]managerHandlerItemWorker, 0, 1)
					} else {
						out.Workers = []managerHandlerItemWorker{}
					}
				} else {
					out.Workers = (out.Workers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 managerHandlerItemWorker
					(v1).UnmarshalEasyJSON(in)
					out.Workers = append(out.Workers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]managerHandlerItemTask, 0, 1)
					} else {
						out.Tasks = []managerHandlerItemTask{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v2 managerHandlerItemTask
					(v2).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "listeners":
			if in.IsNull() {
				in.Skip()
				out.Listeners = nil
			} else {
				in.Delim('[')
				if out.Listeners == nil {
					if !in.IsDelim(']') {
						out.Listeners = make([]managerHandlerItemListener, 0, 2)
					} else {
						out.Listeners = []managerHandlerItemListener{}
					}
				} else {
					out.Listeners = (out.Listeners)[:0]
				}
				for !in.IsDelim(']') {
					var v3 managerHandlerItemListener
					(v3).UnmarshalEasyJSON(in)
					out.Listeners = append(out.Listeners, v3)
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(out *jwriter.Writer, in managerHandlerStats) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tasks_count\":")
	out.Uint64(uint64(in.TasksCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tasks_waiting_count\":")
	out.Uint64(uint64(in.TasksWaitingCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"workers_count\":")
	out.Uint64(uint64(in.WorkersCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"workers_waiting_count\":")
	out.Uint64(uint64(in.WorkersWaitingCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"listeners_count\":")
	out.Uint64(uint64(in.ListenersCount))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"workers\":")
	if in.Workers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v4, v5 := range in.Workers {
			if v4 > 0 {
				out.RawByte(',')
			}
			(v5).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tasks\":")
	if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v6, v7 := range in.Tasks {
			if v6 > 0 {
				out.RawByte(',')
			}
			(v7).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"listeners\":")
	if in.Listeners == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in.Listeners {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerStats) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerStats) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerStats) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerStats) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers3(l, v)
}
func easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(in *jlexer.Lexer, out *managerHandlerResponseSuccess) {
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
func easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(out *jwriter.Writer, in managerHandlerResponseSuccess) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"result\":")
	out.String(string(in.Result))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v managerHandlerResponseSuccess) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd74d837EncodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *managerHandlerResponseSuccess) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd74d837DecodeGithubComKihamoShadowComponentsWorkersInternalHandlers4(l, v)
}
