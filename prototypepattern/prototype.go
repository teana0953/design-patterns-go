package prototypepattern

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	data int
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		data: c.data,
	}
}

func (c *ConcretePrototype) Data() int {
	return c.data
}

func NewConcretePrototype(data int) *ConcretePrototype {
	// complicated initialization process
	return &ConcretePrototype{
		data: data,
	}
}
