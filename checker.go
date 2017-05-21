package overloading

type Checker struct {
	rules []*Rule
}

func NewChecker(rules ...*Rule) *Checker {
	c := new(Checker)
	c.rules = make([]*Rule, len(rules))

	for i, r := range rules {
		c.rules[i] = r
	}

	return c
}

func (c *Checker) Check(args []interface{}) []interface{} {
	for _, r := range c.rules {
		out, err := r.Check(args)
		if err == nil {
			return out
		}
	}

	panic("No valid rule")
}
