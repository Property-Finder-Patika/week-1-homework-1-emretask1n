// --Notes--

// A struct is a type for data abstraction in Go.

// Functions can be associated with struct types.
// >> There is no class or equivalent type in Go!

//Function greet() is associated with struct Person and thus becomes a method.

// After created instances of Person, greet() method can be called on them.

package main

import (
	"fmt"
)

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + g.name + " :)"
}
func main() {
	greeter := Person{"Nihal"}
	var greeting = greeter.greet()
	fmt.Printf("%s\n", greeting)
}
