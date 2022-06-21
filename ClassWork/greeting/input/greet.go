// --Notes--

//An input can be received from the user using Scanln function of fmt.

package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}
func createGreet(name string) string {
	return "Selam " + name + " :)"
}
