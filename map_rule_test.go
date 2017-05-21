package overloading

import "testing"

func TestMapRule(t *testing.T) {
	t.Parallel()
	r := NewMapRule(
		"x", NewArgument("int"),
		"y", NewArgument("int", 3))

	f := func(arg map[string]interface{}) int {
		out, err := r.Check(arg)
		if err != nil {
			t.Error("The argument should be correct")
		}

		na := out["x"].(int)
		nb := out["y"].(int)

		return na + nb
	}

	args := make(map[string]interface{})
	args["x"] = 3

	n := f(args)
	if n != 6 {
		t.Error("Wrong rule")
	}
}

func TestMapRuleViolation(t *testing.T) {
	t.Parallel()

	r := NewMapRule(
		"x", NewArgument("int"),
		"y", NewArgument("int", 3))

	f := func(arg map[string]interface{}) int {
		out, err := r.Check(arg)
		if err == nil {
			t.Error("The argument should be wrong")
		} else {
			return 0
		}

		na := out["x"].(int)
		nb := out["y"].(int)

		return na + nb
	}

	args := make(map[string]interface{})
	args["x"] = 3
	args["z"] = 5

	_ = f(args)
}
