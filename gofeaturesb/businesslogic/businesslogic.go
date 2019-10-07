/*
Package businesslogic is an initial exanple of how to use the Go Testing Framework
*/
package businesslogic

import (
	"errors"
	"math"
)

// DivAndRemainder tests if the argument is -ve it throws an error
func DivAndRemainder(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	d := int(math.Floor(float64(a) / float64(b)))
	r := a % b
	return d, r, nil

}
