# Floating Point Numbers

Go provides two sizes of floating-point numbers, ```float32``` and ```float64```. Their arithmetic
properties are governed by the IEEE 754 standard implemented by all modern CPUs.

A ```float32``` provides approximately six decimal digits of precision, whereas a ```float64```
provides about 15 digits; float64 should be preferred for most purposes because ```float32```
computations accumulate error rapidly unless one is quite careful, and the smallest positive
integer that cannot be exactly represent ed as a ```float32``` is not large:

```go
var f float32 = 16777216 // 1 << 24
fmt.Println(f == f+1) // "true"!
```

Floating-point numbers can be written literally using decimals, like this:

```
const e = 2.71828 // (approximately)
```

Floating-point values are conveniently printed with Printf’s %g verb, which chooses the most
compact representation that has adequate precision, but for tables of data, the %e (exp onent)
or %f (no exponent) for ms may be more appropriate. 

```go
for x := 0; x < 8; x++ {
fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
}
```

The code above prints the pow ers of e with thre e de cimal digits of pre cision, aligne d in an
eig ht-charac ter field:

```go
x=0 eA = 1.000
x=1 eA = 2.718
x=2 eA = 7.389
x=3 eA = 20.086
x=4 eA = 54.598
x=5 eA = 148.413
x=6 eA = 403.429
x=7 eA = 1096.633
```

In addition to a large collection of the usual mathematical functions, the ```math``` package has
functions for creating and detecting the special values defined by IEEE 754: the positive and
negative infinities, which represent numbers of excessive magnitude and the result of division
by zero; and NaN (‘‘not a number’’), the result of such mathematically dubious operations as
0/0 or Sqrt(-1).

```go
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
```

The function math.IsNaN tests whether its argument is a not-a-number value, and math.NaN
returns such a value. It’s tempting to use NaN as a sentinel value in a numeric computation,
but testing whether a specific computational result is equal to NaN is fraught with peril
because any comparison with NaN always yields false:

```go
nan := math.NaN()
fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
```

