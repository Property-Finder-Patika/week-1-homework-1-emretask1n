# Complex Numbers

Go provides two sizes of complex numbers, ```complex64``` and ```complex128```, whose components
are ```float32``` and ```float64``` respectively. The built-in function complex creates a complex number from its real and imaginary 
components, and the built-in real and imag functions extract those components:

```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y) // "(-5+10i)"
fmt.Println(real(x*y)) // "-5"
fmt.Println(imag(x*y)) // "10"
```

Under the rules for constant arithmetic, complex constants can be added to other constants
(integer or floating point, real or imaginary), allowing us to write complex numbers natural ly,
li ke 1+2i, or equivalently, 2i+1. The declarat ions of x and y ab ove can be simplified:

```
x := 1 + 2i
y := 3 + 4i
```

Complex numbers may be compared for equality with == and !=.

The math/cmplx package provides library functions for working with complex numbers,

```
fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
```


