# Overloading

Overloaded Argument Checker for Go

## Intro

Currently, Go lacks function (or method) overloading. One possible solution may be varargs:

```
func Foo(args ...interface{}) interface{} {
    // Process args
    
    // Do somethings
}
```

However, function arguments in this pattern loses the benefit from type checking since these functions use arrays for anything. Argument checking becomes tedious routines in these functions. In our package, we built an argument checker, reducing these chores.

Using varargs for function overloading is still debatable. Don't abuse it.

Another solution may use map as argument:

```
func Bar(arg map[string]interface{}) interface{} {
	// Process arg
	
	// Do something
}
```

Similiarly, our package implements a map argument checker.

## Usage

This ``add`` function responds to two different data types:

```
package main

import (
	ov "github.com/cwchentw/overloading-golang"
	"log"
)

var checker *ov.ListChecker

func init() {
	// Expect (int, int) or (float64, float64)
	checker = ov.NewListChecker(
		ov.NewListRule(
			ov.NewArgument("int"),
			ov.NewArgument("int")),
		ov.NewListRule(
			ov.NewArgument("float64"),
			ov.NewArgument("float64")))
}

func main() {
	n := add(3, 2).(int)

	if n != 5 {
		log.Fatal("Wrong number")
	}
}

func add(args ...interface{}) interface{} {
	out := checker.Check(args)

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
```

This ``add`` function has optional argument:

```
package main

import (
	ov "github.com/cwchentw/overloading-golang"
	"log"
)

var checker *ov.ListChecker

func init() {
	// Expect (int, int) or (int)
	checker = ov.NewListChecker(
		ov.NewListRule(
			ov.NewArgument("int"),
			ov.NewArgument("int", 3)))
}

func main() {
	n := add(3)

	if n != 6 {
		log.Fatal("Wrong number")
	}
}

func add(args ...interface{}) int {
	out := checker.Check(args)

	a := out[0].(int)
	b := out[1].(int)

	return a + b
}

```

Use map as argument:

```
package main

import (
	ov "github.com/cwchentw/overloading-golang"
	"log"
	"reflect"
)

var checker *ov.MapChecker

func init() {
	checker = ov.NewMapChecker(
		ov.NewMapRule(
			"x", ov.NewArgument("int"),
			"y", ov.NewArgument("int")),
		ov.NewMapRule(
			"x", ov.NewArgument("float64"),
			"y", ov.NewArgument("float64")))
}

func main() {
	arg := make(map[string]interface{})
	arg["x"] = 3.0
	arg["y"] = 2.0
	n := add(arg).(float64)
	if n != 5.0 {
		log.Fatal("Wrong checker")
	}
}

func add(arg map[string]interface{}) interface{} {
	out := checker.Check(arg)

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
```

## Copyright

2017, Michael Chen

## License

MIT