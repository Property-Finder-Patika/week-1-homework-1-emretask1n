// --Notes--

//Functions are declared using func keyword.

//Execution always starts by default with the main function which is declared in the main package.

//fmt library is imported using import declaration.

//Println function prints the string argument into the console.

//go build greet.go compiles, links, creates and saves an executable file called greet for further use.

//All Go source files must start with a package declaration and then may have an import declaration to include other packages.

//main package is special; it defines a standalone executable program and it must include a special function, main that starts the execution.

package main

import "fmt"

func main() {
	fmt.Println("Selam :)")

}
