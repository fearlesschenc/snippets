// Package memory 包含内存相关
// * 用什么类型的 receiver 相当于用什么类型的参数
// * 用实例类型去用 receiver 是指针类型的方法会转成指针，反之亦然
// * 把实例类型直接复制不会自动分配内存，仅仅将
// https://golang.org/doc/faq#methods_on_values_or_pointers
package memory

import "fmt"

// Bar
type Bar struct {
	abc int
}

func (b *Bar) print() {
	fmt.Printf("bar is %d", b.abc)
}

func (b *Bar) set(a int) {
	b.abc = a
}

func (b Bar) seta(a int) {
	bb := &b
	_ = bb
	b.abc = a
}

func set(b Bar, a int) {
	b.abc = a
}

type Foo struct {
	m map[string]string
	b Bar
}

func Assign() {
	f1 := Foo{
		m: map[string]string{
			"1": "2",
		},
		b: Bar{abc: 10},
	}
	f2 := Foo{}

	f2 = f1
	_ = f2
}
