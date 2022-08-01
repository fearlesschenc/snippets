package generic

import (
	"fmt"
)

type Model[T any] struct {
	Total int
	Items []T `json:"items"`
}

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Stringer interface {
	String() string
}

func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return
}

type Vector[T AnyString] []T

type MyString string

func (s MyString) String() string {
	return string(s)
}

type AnyString interface {
	~string
}

func Run() {
	//m := map[string]string{
	//	"foo": "bar",
	//}
	//getKeys[string, string](m)
	//Print[string]([]string{"foo", "bar"})

	var vec = Vector{"foo", "bar"}

	for _, v := range vec {
		fmt.Println(v)
	}
}
