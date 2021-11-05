package types

import "fmt"

type Object interface {
	GetName() string
}

type Foo interface {
	Object
}

func NewFoo() Foo {
	return &foo{Name: "foo"}
}

type foo struct {
	Name string
}

func (f *foo) GetName() string {
	return f.Name
}

type Bar interface {
	Object
}

func NewBar() Bar {
	return &bar{Name: "bar"}
}

type bar struct {
	Name string
}

func (b *bar) GetName() string {
	return b.Name
}

func Main() {
	var o Object

	o = NewFoo()
	switch o.(type) {
	case Foo:
		fmt.Println("foo")
	case Bar:
		fmt.Println("bar")
	default:
		fmt.Println("default")
	}
}
