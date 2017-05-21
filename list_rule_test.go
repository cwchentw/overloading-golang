package overloading

import "testing"

func TestRule(t *testing.T) {
	// Expect (int, int)
	r := NewListRule(
		NewArgument("int"),
		NewArgument("int"))

	f := func(args ...interface{}) int {
		out, err := r.Check(args)
		if err != nil {
			t.Log(err)
		}
		a := out[0].(int)
		b := out[1].(int)

		return a + b
	}

	if f(3, 2) != 5 {
		t.Error("Wrong rule")
	}
}

func TestOptionalRule(t *testing.T) {
	// Expect (int, int) or (int)
	r := NewListRule(
		NewArgument("int"),
		NewArgument("int", 3))

	f := func(args ...interface{}) int {
		out, err := r.Check(args)
		if err != nil {
			t.Log(err)
		}
		a := out[0].(int)
		b := out[1].(int)

		return a + b
	}

	if f(3) != 6 {
		t.Error("Wrong rule")
	}
}

func TestRuleViolation(t *testing.T) {
	// Expect (int, int)
	r := NewListRule(
		NewArgument("int"),
		NewArgument("int"))

	f := func(args ...interface{}) int {
		out, err := r.Check(args)
		if err == nil {
			t.Error("There should be some error")
		} else {
			return 0
		}

		a := out[0].(int)
		b := out[1].(int)

		return a + b
	}

	_ = f(3, 4, 5)
}
