package businesslogic

// Here's a basic interface for geometric 3d shapes, with two methods: "volume", "surfacearea"

// For our example we'll implement this interface on
// `sphere` and `cuboid` types.

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `sphere`s.

// The implementation for `cuboid`s.

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

func Example() {
	s := sphere{radius: 4.0}
	c := cuboid{edge: 5.0}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	volume(s)
	volume(c)
	surfaceArea(s)
	surfaceArea(c)

	// Output:
}
