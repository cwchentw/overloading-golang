package overloading

type Argument struct {
	_type      string
	isOptional bool
	_default   interface{}
}

func NewArgument(args ...interface{}) *Argument {
	/* Expected arguments:

	   - (string)
	   - (string, interface{})
	*/
	arr := make([]interface{}, 0)
	for _, e := range args {
		arr = append(arr, e)
	}
	if len(arr) != 1 && len(arr) != 2 {
		panic("Wrong length of arguments")
	}

	p := new(Argument)
	p._type = arr[0].(string)

	if len(arr) == 2 {
		p.isOptional = true
		p._default = arr[1]
	}

	return p
}

func (a *Argument) Type() string {
	return a._type
}

func (a *Argument) IsOptional() bool {
	return a.isOptional
}

func (a *Argument) Default() interface{} {
	return a._default
}
