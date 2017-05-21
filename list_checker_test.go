package overloading

import (
	"fmt"
	"testing"
)

func TestChecker(t *testing.T) {
	t.Parallel()

	// Expect (int, int)
	c := NewListChecker(
		NewRule(
			NewArgument("int"),
			NewArgument("int")))

	f := func(args ...interface{}) int {
		out := c.Check(args)

		a := out[0].(int)
		b := out[1].(int)

		return a + b
	}

	if f(3, 2) != 5 {
		t.Error("Wrong checker")
	}
}

func TestCheckerViolation(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	// Expect (int, int)
	c := NewListChecker(
		NewRule(
			NewArgument("int"),
			NewArgument("int")))

	f := func(args ...interface{}) int {
		out := c.Check(args)

		a := out[0].(int)
		b := out[1].(int)

		return a + b
	}

	_ = f(2, 3, 4)
}

func TestMultiRuleChecker(t *testing.T) {
	t.Parallel()

	// Expect (int, int) or (float64, float64)
	c := NewListChecker(
		NewRule(
			NewArgument("int"),
			NewArgument("int")),
		NewRule(
			NewArgument("float64"),
			NewArgument("float64")))

	f := func(args ...interface{}) interface{} {
		out := c.Check(args)

		switch out[0].(type) {
		case int:
			na := out[0].(int)
			nb := out[1].(int)
			return na + nb
		case float64:
			na := out[0].(float64)
			nb := out[1].(float64)
			return na + nb
		default:
			panic("Unknown type")
		}
	}

	if fmt.Sprint(f(3, 2)) != "5" {
		t.Log(fmt.Sprint(f(3, 2)))
		t.Error("Wrong checker")
	}

	if fmt.Sprint(f(3.0, 2.0)) != "5" {
		t.Log(fmt.Sprint(f(3.0, 2.0)))
		t.Error("Wrong checker")
	}
}
