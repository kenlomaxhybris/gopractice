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
func TestF(t *testing.T) {
	// run with: go test -v -cover
	// also try debugging
	b := F()
	check(b, t)
}

func ExampleF() {
	// This is where the PO/Dev demo the behaviour they expect
	//   run with: go test -v -cover
	//   run with: godoc -http ":8081"
	fmt.Print(F())
	// Output:
	// true
}

func BenchmarkF(b *testing.B) {
	// go test -bench=.
	// gobenchui -last 10 .
	for i := 0; i < b.N; i++ {
		F()
	}
}
