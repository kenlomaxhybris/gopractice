package businesslogic

import "fmt"

func Example() {

	_, e := F(1)
	fmt.Println(e)
	_, e = F(0)
	fmt.Println(e)

	// Output:
	// <nil>
	// No no no!
}
