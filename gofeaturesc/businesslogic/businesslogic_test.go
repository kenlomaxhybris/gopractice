package businesslogic

import (
	"testing"
)

func compare(want, got string, t *testing.T) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %s but got %s", want, got)
	}
}
func TestGreetings(t *testing.T) {
	// run with: go test -v -cover
	t.Run("WW", func(t *testing.T) {
		g := Greetings(America, "Walter", "White")
		compare("Howdi Walter White!", g, t)
	})

	t.Run("DT", func(t *testing.T) {
		g := Greetings(America, "Donald", "Trump")
		compare("Up yours Donald Trump!", g, t)
	})

	t.Run("BJ", func(t *testing.T) {
		g := Greetings(England, "Alexander", "Boris", "de Pfeffel", "Johnson")
		compare("Greetings Alexander Boris de Pfeffel Johnson!", g, t)
	})
}

func ExampleGreetings() {
	// See doc with: godoc -http ":8081"

	// Output:
	// Howdi Walter White!
	// Up yours Donald Trump!
	// Greetings Alexander Boris de Pfeffel Johnson!
	// Bonsieur Emmanuel Macron!
}

func BenchmarkGreetings(b *testing.B) {
	// go test -bench=.
	for i := 0; i < b.N; i++ {
		Greetings(France, "Pierre")
	}
}
