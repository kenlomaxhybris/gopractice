/*
Package businesslogic is an initial exanple of how to use the Go Testing Framework
*/
package businesslogic

// Country is enum for the countries
type Country int

// The possible countries
const (
	France = iota
	America
	England
)

func greetings(names ...string) string {
	s := ""
	for _, n := range names {
		s = s + " " + n
	}
	return s + "!"
}

// Greetings returns an apporpriate greeting
func Greetings(country int, names ...string) string {
	switch country {
	case France:
		return "Bonsieur" + greetings(names...)
	case America:
		if names[0] == "Donald" {
			return "Up yours" + greetings(names...)
		} else {
			return "Howdi" + greetings(names...)
		}
	default:
		return "Greetings" + greetings(names...)
	}
}
