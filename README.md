# Overloading

Overloaded Argument Checker for Go

## Usage

This ``add`` function responds to two different data types:

```
package main

import (
	ov "github.com/cwchentw/overloading-golang"
	"log"
)

var checker *ov.Checker

func init() {
	// Expect (int, int) or (float64, float64)
	checker = ov.NewChecker(
		ov.NewRule(
			ov.NewArgument("int"),
			ov.NewArgument("int")),
		ov.NewRule(
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

## Intro

Currently, Go lacks function (or method) overloading. One possible solution may be varargs:

```
func Foo(args ...interface{}) {
    // Process args
    
    // Do somethings
}
```

However, function arguments in this pattern loses the benefit from type checking since these functions use arrays for anything. Argument checking becomes tedious routines in these functions. In our package, we built an argument checker, reducing these chores.

Using varargs for function overloading is still debatable. Don't abuse it.

## Copyright

2017, Michael Chen

## License

MIT