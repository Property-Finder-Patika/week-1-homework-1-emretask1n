# Variables
A var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value. Each declaration has the general form

``` go
var name type = expression 
```

If the type is omitted, it is determined by the initializer expression. If the expression is omitted, **the initial value is
the zero value for the type**, which is **0 for numbers, false for boole ans, "" for strings, and nil
for interfaces and reference types (slice, point er, map, channel, function)**. The zero value of an
aggregate typ e li ke an array or a struct has the zero value of all of its elements or fields


```go
var s string
fmt.Println(s) 
// "" is the output
```

It is possible to declare and optionally initialize a set of variables in a single declaration, with a
matching list of expressions. Omitting the type allows declaration of multiple variables of different types:

```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

### Short Variable Declarations

It takes the form **name := expression**, and the **type of name** is
determined by the **type of expression**. Here are three of the many short variable declarations
in the lissajous function (§1.4):

```go
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64() * 3.0
t := 0.0
```

Because of their brevity and flexibility, short variable declarations are used to declare and initialize the majority of local variables

As wit h var de clarations, multiple variables may be declared and initialized in the same short
variable declaration,

```go
i, j := 0, 1
```

declarations with multiple initializer expressions should be used only when they help readability, such as for short and natural groupings like the initialization part of a for loop.

**Keep in mind that := is a declaration, whereas = is an assignment.**

Like ordinary var declarations, short variable declarations may be used for cal ls to functions
li ke os.Open that return two or more values:

```go
f, err := os.Open(name)
if err != nil {
return err
}
// ...use f...
f.Close()
```

In the code below, the first statement declares both in and err. The second declares out but
only assigns a value to the existing err variable

```go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```

A short variable declaration must declare at least one new variable, however, so this code will
not compile

```go
, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables
```

The fix is to use an ordinary assignment for the second statement.


### Pointers

A variable is a **piece of storage** containing a value.

**A pointer value is the address of a variable.** A pointer is thus the location at which a value is
stored. Not every value has an address, but every variable does. With a pointer, we can read
or update the value of a variable **indirectly.**

If a variable is declared var x int, the expression &x (‘‘address of x’’) yields a pointer to an
integer variable, that is, a value of type *int, which is pronounced ‘‘pointer to int.’’ If this value is called p,
we say ‘‘p points to x,’’ or equivalently ‘‘p contains the address of x.’’ The variable to which p points is written *p. The expression *p yields 
the value of that variable, an int, but since *p denotes a variable, it may also appear on the left-hand side of an assignment,
in which cas e the assignment updates the variable.


```go
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2 // equivalent to x = 2
fmt.Println(x) // "2"
```

**The zero value for a pointer of any type is nil.** The test p != nil is true if p points to a variable. Pointers 
are comparable; two pointers are equal if and only if the y point to the same
variable or both are nil.

```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```


### The new function

Another way to create a variable is to use the built-in function new. The expression new(T)
creates an unnamed variable of type T, initializes it to the zero value of T, and returns its
address, which is a value of type *T.

```go
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // "0"
*p = 2 // sets the unnamed int to 2
fmt.Println(*p) // "2"
```


### Lifetime Of Variables

The lifetime of a variable is the interval of time during which it exists as the program executes.
The lifetime of a package-level variable is the entire execution of the program.