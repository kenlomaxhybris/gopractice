package businesslogic

import (
	"fmt"
)

func ExampleFor() {
	// run with: go test -v -cover
	// See doc with: godoc -http ":8081"

	for i := 0; i < 100; i++ {
		if j := i * 2; j == 44 {
			continue
		}
		if i == 99 {
			break

		}
		fmt.Println(i)
	}
	// Output:
	// 1 is an int
	// bob is a string
	// ..

}
