// Boiling prints the boiling point of water.

//- Notes-

// The constant boilingF is a package-level declaration, whereas the variables f and
//c are local to the function main. The name of each package-level entity is visible not only
//throughout the source file that contains its declaration, but throughout all the files of the package. By contrast, local declarations are visible only wit hin the function in which they are
//de clare d and perhaps only wit hin a small part of it.

//A function declaration has a name, a list of parameters (the variables whose values are
//provided by the function’s callers), an optional list of results, and the function body, which
//contains the statements that define what the function does. The result list is omitted if the
//function does not return anything.

package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
