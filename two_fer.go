// Exercism :: Exercise #2 :: Two-fer Package
package twofer

import "fmt"

// ShareWith generates a string, "One for {<name>, you}, one for me", according to a name specifier
// and returns the resulting string.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
