package sync

import "sync"

func Pool() {
	// Pool 是一个大的可用对象的池子，新建的时候定义一个 New 函数用来创建
	pool := sync.Pool{New: func() interface{} {
		return new(map[string]string)
	}}

	// Pool 中的对象随时可能会被删，常用来保存在多个 goroutine 中都要使用
	// 的常用对象，以减轻 GC 压力
	m := make(map[string]string)
	m["foo"] = "bar"
	pool.Put(m)

	n := pool.Get().(map[string]string)
	println(n["foo"])
}
