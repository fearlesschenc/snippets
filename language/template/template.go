package template

import (
	"os"
	"text/template"
)

func Main() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, _ := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	tmpl.Execute(os.Stdout, sweaters)
}
