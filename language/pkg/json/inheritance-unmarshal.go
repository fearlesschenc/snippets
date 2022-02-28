package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	pkgstring "golang_snippets/language/pkg/string"
)

type StructB struct {
	pkgstring.StructA

	Name string `json:"name"`
}

func InheritanceUnmarshal() {
	a := &pkgstring.StructA{Age: 99}

	b := &StructB{
		StructA: pkgstring.StructA{
			Age: 100,
		},

		Name: "bb",
	}

	c := &StructB{
		StructA: *a,

		Name: "cc",
	}

	data, _ := json.Marshal(b)
	fmt.Println(string(data))

	data, _ = json.Marshal(c)
	fmt.Println(string(data))

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(b)
	fmt.Println(buf.String())

	aaa := pkgstring.StructA{}
	err := json.Unmarshal([]byte(`{"age": 1}`), aaa)
	if err != nil {
		panic(err)
	}
}
