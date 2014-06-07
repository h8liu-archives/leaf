package scope

import (
	"errors"
)

type Context struct {
	scopes []*Scope
}

func NewContext() *Context {
	ret := new(Context)
	ret.scopes = make([]*Scope, 1, 10)
	ret.scopes[0] = New()

	return ret
}

func (c *Context) Level() int {
	return len(c.scopes)
}

func (c *Context) Top() *Scope {
	return c.scopes[len(c.scopes)-1]
}

func (c *Context) IsGlobal() bool {
	return c.Level() == 1
}

func (c *Context) Define(item Item) error {
	return c.Top().Define(item)
}

func (c *Context) Query(name string) Item {
	n := len(c.scopes)
	for i := range c.scopes {
		s := c.scopes[n-1-i]
		ret := s.Query(name)
		if ret != nil {
			return ret
		}
	}
	return nil
}

func (c *Context) Enter() {
	c.scopes = append(c.scopes, New())
}

func (c *Context) Exit() error {
	if c.IsGlobal() {
		return errors.New("cannot exit global context")
	}

	c.scopes = c.scopes[:len(c.scopes)-1]
	return nil
}
