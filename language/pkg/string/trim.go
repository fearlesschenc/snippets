package string

import (
	"fmt"
	"strings"
)

type StructA struct {
	Age int `json:"age"`
}

func Trim() {
	fmt.Println(strings.TrimLeft(".....abc", "."))
	fmt.Println(strings.TrimRight("abc.....", "."))
}

func Title() {
	fmt.Println(strings.Title("abc"))
}
