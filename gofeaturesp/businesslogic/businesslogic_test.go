package businesslogic

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(),
		"it didn't workkkk",
	}
}

func main() {
	err := run()
	fmt.Println(err)

}
