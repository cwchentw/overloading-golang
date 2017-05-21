package overloading

type ListChecker struct {
	rules []*ListRule
}

func NewListChecker(rules ...*ListRule) *ListChecker {
	c := new(ListChecker)
	c.rules = make([]*ListRule, len(rules))

	for i, r := range rules {
		c.rules[i] = r
	}

	return c
}

func (c *ListChecker) Check(args []interface{}) []interface{} {
	for _, r := range c.rules {
		out, err := r.Check(args)
		if err == nil {
			return out
		}
	}

	panic("No valid rule")
}
