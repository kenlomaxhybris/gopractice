/*
Package businesslogic is an initial exanple of how to use the Go Testing Framework
*/
package businesslogic

import (
	"math/rand"
	"time"
)

// F is a trivial function that returns the value true!
func F() bool {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return true
}
