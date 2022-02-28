package set

import (
	"fmt"

	"github.com/deckarep/golang-set"
)

type structA struct {
	Name string
}

func Set() {
	s := mapset.NewSet()
	s.Add(1)
	s.Add(2)

	ss := mapset.NewSet()
	ss.Add(2)
	ss.Add(3)

	fmt.Println(s.Difference(ss).ToSlice())
}
