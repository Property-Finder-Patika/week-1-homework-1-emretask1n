// --Notes--

// main function can receive arguments supplied through os.Args

package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	greeting := createGreet(name)
	fmt.Printf("%s\n", greeting)
}
func createGreet(name string) string {
	return "Selam " + name + " :)"
}
