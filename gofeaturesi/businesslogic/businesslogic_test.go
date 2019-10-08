package businesslogic

import "fmt"

func myDiv(a float64, b float64) {
	if b == 0 {
		panic("Cannot divide by zero")
	}
}

func Example() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("In recover")
		}
	}()

	fmt.Println("A")
	myDiv(1, 2)
	fmt.Println("B")
	myDiv(1, 0)
	fmt.Println("C")

	// Output:
	// A
	// B
	// In recover
}
