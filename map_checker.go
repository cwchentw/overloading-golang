package overloading

type MapChecker struct {
	rules []*MapRule
}

func NewMapChecker(args ...*MapRule) *MapChecker {
	m := new(MapChecker)
	m.rules = make([]*MapRule, len(args))

	for i, a := range args {
		m.rules[i] = a
	}

	return m
}

func (m *MapChecker) Check(args map[string]interface{}) map[string]interface{} {
	for _, r := range m.rules {
		out, err := r.Check(args)
		if err == nil {
			return out
		}
	}

	panic("No valid argument")
}
