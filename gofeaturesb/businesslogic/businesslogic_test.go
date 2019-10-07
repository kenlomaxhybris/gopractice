package businesslogic

import (
	"fmt"
	"testing"
)

func check(b bool, t *testing.T) {
	t.Helper()
	if !b {
		t.Errorf("Expected true but got %v", b)
	}
}
func TestDivAndRemainder(t *testing.T) {
	// run with: go test -v -cover
	// also try debugging
	t.Run("Happy path", func(t *testing.T) {
		d, r, e := DivAndRemainder(7, 2)
		if d != 3 || r != 1 || e != nil {
			t.Error("Unexpected")
		}
	})
	t.Run("Div by zero", func(t *testing.T) {
		d, r, e := DivAndRemainder(7, 0)
		if d != 0 || r != 0 || e == nil {
			t.Error("Unexpected")
		}
	})
}

func ExampleDivAndRemainder() {
	// run with: go test -v -cover
	// See doc with: godoc -http ":8081"
	fmt.Println(DivAndRemainder(7, 2))
	fmt.Println(DivAndRemainder(7, 0))

	// Output:
	// 3 1 <nil>
	// 0 0 Cannot divide by zero

}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivAndRemainder(1<<10, 12345)
	}
}
