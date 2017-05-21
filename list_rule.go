package overloading

import (
	"errors"
	"reflect"
)

type ListRule struct {
	args []*Argument
	op   int
}

func NewListRule(args ...*Argument) *ListRule {
	r := new(ListRule)
	r.args = make([]*Argument, len(args))

	for i, a := range args {
		r.args[i] = a
		if a.IsOptional() {
			r.op += 1
		}
	}

	return r
}

func (r *ListRule) Check(args []interface{}) ([]interface{}, error) {
	_len := len(r.args)
	out := make([]interface{}, _len)

	n := 0
	for i, e := range args {
		if i >= _len {
			return out, errors.New("Argument length is too long")
		}
		t := reflect.TypeOf(e).String()
		if t != r.args[i].Type() {
			return out, errors.New("Wrong argument type " + t + " at " + string(i))
		}
		out[i] = e
		n++
	}

	if n < _len-r.op || _len < n {
		return out, errors.New("Wrong argument length")
	}

	for i := n; i < _len; i++ {
		out[i] = r.args[i].Default()
	}

	return out, nil
}
