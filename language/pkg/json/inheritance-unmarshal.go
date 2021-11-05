package json

import (
	"encoding/json"
	"fmt"
)

type StructA struct {
	Age int `json:"age"`
}

type StructB struct {
	StructA
}

func inheritanceUnmarshal() {
	b := StructB{}

	bytes := []byte(`{"age": 10000}`)
	json.Unmarshal(bytes, &b)
	fmt.Println(b.Age)
}
