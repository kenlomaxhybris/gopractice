package businesslogic

import "fmt"

type MyError struct {
}

func (e MyError) Error() string {
	return fmt.Sprintf("No no no!")
}

func F(n int) (int, error) {
	if n == 0 {
		return 0, MyError{}
	}
	return n, nil
}
