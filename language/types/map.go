package types

import "fmt"

type empty struct{}

type request struct {
	Namespace string
	Name      string
}

// MapWithInterface 展示如何使用 map 存储非常规 Key
// map 中可以存放任何可以比较的类型，具体可以参考(https://golang.org/ref/spec#Comparison_operators)
// interface 的 Key 比较的是对应的动态类型(潜在类型)
func MapWithInterface() {
	m := make(map[interface{}]empty)
	m[request{
		Namespace: "foo",
		Name:      "bar",
	}] = empty{}
	m[request{
		Namespace: "foo",
		Name:      "bar",
	}] = empty{}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(len(m))
}
