# Constants

Constants are expressions whose value is known to the compiler and whose evaluation is guaranteed to occur at compile time,
not at runtime. The underlying type of every constant is a basic type: boolean, string, or number

a constant is more appropriate than a variable for a mathematical cons tant like
pi, since its value won’t change:

```go
const pi = 3.14159 // approximately; math.Pi is a better approximation
```

```go
const (
e = 2.71828182845904523536028747135266249775724709369995957496696763
pi = 3.14159265358979323846264338327950288419716939937510582097494459
)
```

In the following , time.Duration is a named type whose underlying type is int64,
and time.Minute is a const ant of that type. Both of the constants declared below thus have the type time.Duration 
as well, as revealed by %T:

```go
const noDelay time.Duration = 0
const timeout = 5 * time.Minute
fmt.Printf("%T %[1]v\n", noDelay) // "time.Duration 0"
fmt.Printf("%T %[1]v\n", timeout) // "time.Duration 5m0s
fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
```


### 3.6.1. The Constant Generator iota

A const declaration may use the constant generator iota, which is used to create a sequence
of related values without spelling out each one explicitly. In a const declaration, the value of
iota begins at zero and increments by one for each item in the sequence.

```go
type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

This declares Sunday to be 0, Monday to be 1, and so on.

### 3.6.2. Untyped Constants

Constants in Go are a bit unusual. Although a constant can have any of the basic data types
like int or float64, including named basic types like time. Duration, many constants are
not committed to a particular type. The compiler represents these uncommitted constants
with much greater numeric precision than values of basic types, and arithmetic on them is
more precise than machine arithmetic; you may assume at least 256 bits of precision. There
are six flavors of these uncommitted constants, called untyped boolean, untyped integer,
untyped rune, untyped floating-point, untyped complex, and untyped string

```go
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi
```

If math.Pi had been committed to a specific type such as float64, the result would not be as
precise, and type conversions would be required to use it when a float32 or complex128
value is wanted:

```go
const Pi64 float64 = math.Pi
var x float32 = float32(Pi64)
var y float64 = Pi64
var z complex128 = complex128(Pi64)
```

Whether implicit or explicit, converting a constant from one type to another requires that the
target type can represent the original value. Rounding is allowed for real and complex floating-point numbers:

```go
const (
deadbeef = 0xdeadbeef // untyped int with value 3735928559
a=uint32(deadbeef) // uint32 with value 3735928559
b=float32(deadbeef) // float32 with value 3735928576 (rounded up)
c=float64(deadbeef) // float64 with value 3735928559 (exact)
d=int32(deadbeef) // compile error: constant overflows int32
e=float64(1e309) // compile error: constant overflows float64
f=uint(-1) // compile error: constant underflows uint
)
```
