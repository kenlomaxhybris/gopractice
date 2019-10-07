/*
Package businesslogic is an initial exanple of how to use the Go Testing Framework
*/
package businesslogic

import "fmt"

// WhatAmI describves the object sent in
func WhatAmI(o interface{}) string {
	return fmt.Sprintf("%T, %#v, %v", o, o, o)
}
