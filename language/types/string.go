package types

import "fmt"

func IterateString() {
	sample := "01234567"
	output := []rune(sample)
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
		output[i] -= 1
		//fmt.Println(reflect.TypeOf(output[i]))
	}
	fmt.Println(string(output))
}
