# Assignments

The value held by a variable is updated by an assignment statement, which in its simplest form
has a variable on the left of the = sign and an expression on the right

```go
x=1 // named variable
*p = true // indirect variable
person.name = "bob" // struct field
count[x] = count[x] * scale // array or slice or map element
```

### Tuple Assignment

Another form of assignment, known as tuple assignment, allows several variables to be
assigned at once. 

```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```

Tuple assignment can also make a sequence of trivial assignments more compact,

```go
i, j, k = 2, 3, 5
```

When such a call is used in an assignment statement, the left-hand side must have as many
variables as the function has results.

```go
f, err = os.Open("foo.txt") // function call returns two values
```

As with variable declarations, we can assign unwanted values to the blank identifier.

```go
_, err = io.Copy(dst, src) // discard byte count
_, ok = x.(T) // check type but discard result
```


### Assignability


