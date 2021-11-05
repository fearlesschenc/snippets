package goroutine

import (
	"fmt"
	"sync"
)

func WaitGroup() {
	var wg sync.WaitGroup

	values := []int{1, 2, 3, 4, 5}

	for _, val := range values {
		go func(val interface{}) {
			wg.Add(1)
			defer wg.Done()

			fmt.Println(val)
		}(val)
	}

	wg.Wait()
	fmt.Println("Done")
}

func ClosedGoRoutine() {
	fmt.Println("ClosedGoRoutine")

	var once sync.Once
	var wg sync.WaitGroup

	errChan := make(chan error)

	newStuff := func() {
		once.Do(func() {
			fmt.Println("new stuff")
			close(errChan)
		})

		fmt.Printf("%v\n", <-errChan)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			newStuff()
		}()
	}

	wg.Wait()
}

func Main() {
	//WaitGroup()

	ClosedGoRoutine()
}
