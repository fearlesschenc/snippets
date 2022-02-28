package language

import (
	"fmt"
	"sync"
	"time"
)

func ForLoop() {
	a := 1

Loop:
	for j := 0; j < 3; j++ {
		for a < 10 {
			fmt.Println(a)
			if a == 5 {
				continue Loop
			}
			a++
		}
	}

	fmt.Println(a)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		count := 0
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Println("abc")

				if count == 2 {
					return
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("done")
}
