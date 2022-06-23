# Booleans

A value of type bool, or boolean, has only two possible values, true and false. The conditions in if and for statements are booleans, and comparison operators like == and < produce
a boolean result. The unary operator ! is logical negation, so ```!true``` is ```false```, or, one might
say, ```(!true==false)==true```, although as a matter of style, we always simplify redundant
boolean expressions like ```x==true``` to ```x```.

Boolean values can be combined with the && (AND) and || (OR) operators, which have short circuit behavior: if the answer is already determined by the 
value of the left operand, the right operand is not evaluated, making it safe to write expressions like this:

```go
s != "" && s[0] == 'x'
```

The inverse operation is so simple that it doesn’t warrant a function, but for symmetry here it
is:

```go
// itob reports whether i is non-zero.
func itob(i int) bool { return i != 0 }
```
