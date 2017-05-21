package overloading

import (
	"reflect"
	"testing"
)

func TestMapChecker(t *testing.T) {
	c := NewMapChecker(
		NewMapRule(
			"x", NewArgument("int"),
			"y", NewArgument("int")),
		NewMapRule(
			"x", NewArgument("float64"),
			"y", NewArgument("float64")))

	f := func(arg map[string]interface{}) interface{} {
		out := c.Check(arg)

		t := reflect.TypeOf(out["x"]).String()
		switch t {
		case "int":
			na := out["x"].(int)
			nb := out["y"].(int)
			return na + nb
		case "float64":
			na := out["x"].(float64)
			nb := out["y"].(float64)
			return na + nb
		default:
			panic("Unknown type")
		}
	}

	arg1 := make(map[string]interface{})
	arg1["x"] = 3
	arg1["y"] = 2
	if f(arg1).(int) != 5 {
		t.Error("Wrong map checker")
	}

	arg2 := make(map[string]interface{})
	arg2["x"] = 3.0
	arg2["y"] = 2.0
	if f(arg2).(float64) != 5.0 {
		t.Error("Wrong map checker")
	}
}
