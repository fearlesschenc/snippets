package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TypeA struct {
	Age int `json:"age"`
}

type TypeB struct {
	Name string `json:"name"`
}

type Foo struct {
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

var data = []byte(`
{
	"type": "a",
	"config": {
		"age": 10
	}
}
`)

func dynamicUnmarshal() {
	var foo Foo
	json.Unmarshal(data, &foo)
	switch foo.Type {
	case "a":
		v := reflect.TypeOf(foo.Config)
		fmt.Println(v.Name())
		typeA, ok := foo.Config.(*TypeA)
		if !ok {
			panic("invalid typeA")
		}

		fmt.Println(typeA.Age)
	case "b":
		typeB, ok := foo.Config.(*TypeB)
		if !ok {
			panic("invalid typeA")
		}

		fmt.Println(typeB.Name)
	}
}
