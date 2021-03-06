package internal

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/go-gorp/gorp"
	"github.com/kihamo/gotypes"
)

type StructType map[string]interface{}

type TypeConverter struct{}

func (o StructType) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})

	for i := range o {
		data[gotypes.ToString(i)] = o[i]
	}

	return json.Marshal(data)
}

func (o *StructType) UnmarshalJSON(data []byte) error {
	target := make(map[string]interface{})

	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	if target == nil {
		return &json.UnmarshalTypeError{Value: "null", Type: reflect.TypeOf(o)}
	}

	*o = StructType{}
	for i := range target {
		(*o)[i] = target[i]
	}

	return nil
}

func (t TypeConverter) ToDb(val interface{}) (interface{}, error) {
	if t, ok := val.(StructType); ok {
		b, err := json.Marshal(t)

		if err != nil {
			return "", err
		}

		return string(b), nil
	}

	return val, nil
}

func (t TypeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	if _, ok := target.(*StructType); ok {
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)

			if !ok {
				return errors.New("FromDb: Unable to convert JsonField to *string")
			}

			return json.Unmarshal([]byte(*s), target)
		}

		return gorp.CustomScanner{
			Holder: new(string),
			Target: target,
			Binder: binder,
		}, true
	}

	return gorp.CustomScanner{}, false
}
