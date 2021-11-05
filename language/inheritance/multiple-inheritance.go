package inheritance

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type StatusClient interface {
	Status()
}

type Foo struct {
}

func (f *Foo) Read() {

}

func (f *Foo) Status() {

}

type Bar struct {
}

func (b *Bar) Write() {

}

type FooBar struct {
	Reader
	Writer
	StatusClient
}

type Client interface {
	Reader
	Writer
	StatusClient
}

var _ Client = &FooBar{
	Reader: &Foo{},
	Writer: &Bar{},
}
