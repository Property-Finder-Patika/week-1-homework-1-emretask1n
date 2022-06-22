# 2.6 Packages and Files

Packages in Go serve the same purposes as libraries or modules in other languages, supporting
**modularity, encapsulation, separate compilation, and reuse**.

Each file starts wit h a package de claration that defines the package name

When the package is imported, its members are referred to as tempconv.CToF and so on. Package-level names
like the types and constants declared in one file of a package are visible to all the other files of
the package, as if the source code were all in a single file. **Note that tempconv.go imports fmt,
but conv.go does not, because it does not use anything from fmt.**

#### [Exercise 2.1](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter2/packages-and-files/tempconv.go)
Add types, constants, and functions to ```tempconv``` for processing temperatures in the
Kelvin scale, where ```zero``` Kelvin is -273.15℃ and a difference of 1K has the same magnitude as 1℃.


### 2.6.1 Imports 

Within a Go program, every package is identified by a unique string called its import path.
These are the strings that appear in an import declaration like "gopl.io/ch2/tempconv"

In addition to its import path, each package has a package name, which is the short 
(and not necessarily unique) name that appears in its package declaration.


### 2.6.2 Package Initialization

Package initialization begins by initializing package-level variables in the order in which they
are declared, except that dependencies are resolved first:

```go
var a = b + c // a initialized third, to 3
var b = f() // b initialized second, to 2, by calling f
var c = 1 // c initialized first, to 1
func f() int { return c + 1 }
```

If the package has multiple .go files, they are initialized in the order in which the files are
given to the compiler ; the go tool sorts .go files by name before invoking the compiler.
