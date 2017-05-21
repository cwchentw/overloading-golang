package overloading

import (
	"errors"
	"reflect"
)

type MapRule struct {
	m map[string]*Argument
}

func NewMapRule(args ...interface{}) *MapRule {
	_len := len(args)
	if _len%2 != 0 {
		panic("Wrong argument length")
	}
	r := new(MapRule)
	r.m = make(map[string]*Argument)

	for i := 0; i < _len; i += 2 {
		k := args[i].(string)
		r.m[k] = args[i+1].(*Argument)
	}

	return r
}

func (r *MapRule) Check(args map[string]interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	for k, v := range args {
		arg, ok := r.m[k]
		if !ok {
			return out, errors.New("Unknown argument " + k)
		}

		t := reflect.TypeOf(v).String()
		if t != arg.Type() {
			return out, errors.New("Wrong argument type on " + k)
		}

		out[k] = v
	}

	for k := range r.m {
		_, ok := args[k]
		if !ok {
			if r.m[k].IsOptional() {
				out[k] = r.m[k].Default()
			} else {
				return out, errors.New("No default value on " + k)
			}
		}
	}

	return out, nil
}
