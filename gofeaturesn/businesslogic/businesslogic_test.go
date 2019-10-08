package businesslogic

import (
	"fmt"
	"time"
)

func rx(ch chan int) {
	//for r := range ch {
	for {
		a, more := <-ch
		if more {
			fmt.Printf("Rx %d %t\n", a, more)
		}

	}
}

func Example() {
	ch := make(chan int)

	go rx(ch)

	fmt.Println("A")
	for i := 0; i < 2; i++ {
		ch <- i
	}

	close(ch)
	time.Sleep(time.Second)
	fmt.Println("Done")

	//Output:
	// A
	// Rx 0 true
	// Rx 1 true
	// Done
}
