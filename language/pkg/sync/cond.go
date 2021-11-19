package sync

import (
	"fmt"
	"sync"
	"time"
)

var zooOpen = false

func Cond() {
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 2; i++ {
		// visitor
		go func(index int) {
			// 因为调用 Wait 需要修改 Cond 里的数据，所以要在这边加锁
			cond.L.Lock()
			for !zooOpen {
				fmt.Printf("Wait from %d\n", index)
				// Wait 的过程是先把 goroutine 加入到一个通知队列，然后解锁
				// 解锁完等待被通知（cond.Signal/cond.Broadcast）
				// 收到通知之后，会将锁再锁上
				cond.Wait()
				fmt.Printf("It's Open from %d\n", index)
			}

			fmt.Printf("Hello from %d\n", index)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	zooOpen = true
	cond.Signal()
	time.Sleep(time.Second)
	cond.Signal()
	time.Sleep(time.Second)
}
