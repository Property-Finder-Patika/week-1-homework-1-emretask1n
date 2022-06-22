# 2.7 Scope

A declaration associates a name with a program entity, such as a function or a variable. The
scope of a declaration is the part of the source code where a use of the declared name refers to
that declaration

The scope of a declaration is a region of the program text; it is a compile-time property

```go
func f() {}
var g = "g"
func main() {
f := "f"
fmt.Println(f) // "f"; local var f shadows package-level func f
fmt.Println(g) // "g"; package-level var
fmt.Println(h) // compile error: undefined: h
}
```

Most blocks are created by control-flow constructs like ```if``` statements and
```for``` loops. The program below has three different variables called x because each declaration
appears in a different lexical block.

```go
func main() {
    x := "hello!"
    for i := 0; i < len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
        }
    }
}
```

The example below als o has three variables named x, each declared in a different block—one
in the function body, one in the for statement’s block, and one in the loop body—but only two
of the blocks are explicit

```go
func main() {
    x := "hello"
    for _, x := range x {
        x := x + 'A' - 'a'
        fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
        }
}
```

Like for loops, ```if``` statements and ```switch``` statements also create implicit blocks in addition to
their body blocks. The code in the following ```if-else``` chain shows the scope of x and y:

```go
if x := f(); x == 0 {
    fmt.Println(x)
} else if y := g(x); x == y {
    fmt.Println(x, y)
} else {
    fmt.Println(x, y)
}
fmt.Println(x, y) // compile error: x and y are not visible here
```

the scope of f is just the if statement, so f is not accessible to the statements that follow,
resulting in compiler errors.

```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
    return err
    }
f.ReadByte() // compile error: undefined f
f.Close() // compile error: undefined f
```

Thus it is often necessary to declare f before the condition so that it is accessible after

```go
f, err := os.Open(fname)
if err != nil {
    return err
    }
f.ReadByte()
f.Close()
```

Current Go compilers detect that the local cwd variable is never used and report this as an
error, but they are not strictly required to perform this check. Furthermore, a minor change ,
such as the addition of a logging statement that refers to the local cwd would defeat the check.

```go
var cwd string
func init() {
    cwd, err := os.Getwd() // NOTE: wrong!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}
```

There are a number of ways to deal with this potential problem. The most direct is to avoid :=
by declaring err in a separate var de claration:

```go
var cwd string
func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```