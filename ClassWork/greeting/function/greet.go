// --Notes--

// main function can call other functions

// main function takes no arguments and returns no value.

package main
import (
	"fmt"
)
func main() {
	greet(“Zeynep”)
}
func greet(name string) {
	fmt.Printf("Selam %s :)\n", name)
}