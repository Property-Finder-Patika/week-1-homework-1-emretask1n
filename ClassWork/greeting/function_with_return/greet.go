// --Notes--

// main function can call another function createGreet which receives a parameter of type string and returns back another string.

// Strings can be concatenated by + operator.

// var keyword, an identifier and a type are needed to declare a variable.

// If an initial value is supplied through assignment operator =, type can be omitted.
// >> This is called type inference i.e. type is inferred from the right hand side of the assignment i.e. initial value.
// >> If no initial value is supplied then type must be used.

// If an initial value is supplied and := is used for assignment then var can also be omitted.
// >> This can only be done inside a function, outside var must be there,
// >> := is for declaration with assignment, whereas = is for assignment only.
package main

import (
	"fmt"
)

func main() {
	var name string = "Betül"
	var greeting = createGreet(name)
	fmt.Printf("%s", greeting)
}
func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}
