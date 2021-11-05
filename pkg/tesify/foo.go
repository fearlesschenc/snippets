package tesify

//go:generate mockery --name=Adder
type Adder interface {
	Add(int, int) int
}

func Foo(adder Adder, x, y int) int {
	return adder.Add(x, y)
}
