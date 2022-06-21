// --Notes--

// Go is a block-structured language where blocks are declared using a pair of brackets.

// Go pays lots of attention to code formatting; it also has a special tool, go fmt for this purpose.

// go fmt format.go formats the file according to the rules of the language.
package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + p.name + " :)"
}

func main() {
	var greeter Person = Person{"Nihal"}
	greeting :=
		greeter.greet()
	fmt.Printf("%s\n", greeting)
}
