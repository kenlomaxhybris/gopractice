package businesslogic

import (
	"fmt"
	"testing"
)

func TestWhatAmI(t *testing.T) {
	// run with: go test -v -cover
	// also try debugging

}

func ExampleWhatAmI() {
	// run with: go test -v -cover
	// See doc with: godoc -http ":8081"
	s := make([]int, 5)
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", s, s, len(s), cap(s))

	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", a, a, len(a), cap(a))

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", c, c, len(c), cap(c))

	b := a[1:4]
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", b, b, len(b), cap(b))
	b[1] = 99
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", b, b, len(b), cap(b))
	fmt.Printf("Type: %T Instantiation: %#v len %d cap %d\n", a, a, len(a), cap(a))

	for i, v := range s {
		fmt.Printf("i %d v %d\n", i, v)
	}
	// Output:
	// 1 is int
	// bob is string
	// ..

}
