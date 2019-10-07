package businesslogic

import (
	"fmt"
	"math/cmplx"
	"testing"
)

func check(b bool, t *testing.T) {
	t.Helper()
	if !b {
		t.Errorf("Expected true but got %v", b)
	}
}
func TestF(t *testing.T) {
	// run with: go test -v -cover
	// also try debugging
	t.Error("ToDo")
}

func ExampleGreetings() {
	// run with: go test -v -cover
	// See doc with: godoc -http ":8081"

	a := 1
	s := "Bod"
	f := 3.14
	i := int(f)
	var x int
	var y float64
	var b bool
	var ToBe bool = false
	var MaxInt uint64 = 1<<64 - 1
	var z complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println(WhatAmI(a))
	fmt.Println(WhatAmI(s))
	fmt.Println(WhatAmI(i))

	// Output:
	// 1 is an int
	// bob is a string
	// ..

}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WhatAmI("lajslkjsad")
	}
}
