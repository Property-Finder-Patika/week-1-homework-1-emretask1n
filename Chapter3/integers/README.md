# 3.1 Integers

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of
signed integers—8, 16, 32, and 64 bits—represented by the types int8, int16, int32, and
int64, and corresponding unsigned versions uint8, uint16, uint32, and uint64.

There are also two types called just int and uint that are the natural or most efficient size for
signed and unsigned integers on a particular platform; int is by far the most widely used
numeric type. Both these types have the same size, either 32 or 64 bits, but one must not make
assumptions about which; different compilers may make different choices even on identical
hardware.

Go’s binary operators for arithmetic, logic, and comparison are listed here in order of decreasing precedence:

``` go
*   /    %   <<     >>  &   &^
+   -   |   ^
==  !=   <  <=  >   >=
&&
||
```

Two integers of the same type may be compared using the binary comparison operators
below; the type of a comparison expression is a boolean.

``` go
== equal to
!= not equal to
< less than
<= less than or equal to
> greater than
>= greater than or equal to
```

In fact, all values of basic typ e—booleans, numbers, and strings—are comparable, meaning
that two values of the same type may be compared using the == and != operators.

Go also provides the following bit wise binary operators,
``` go
& bit w ise AND
| bit w ise OR
^ bit w ise XOR
&^ bit cle ar (AND NOT)
<< lef t shif t
>> right shif 
```

The code below shows how bit wise operations can be used to interpret a uint8 value as a
compact and efficient set of 8 independent bits.

``` go
var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2
fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}
fmt.Printf("%08b\n", x&y) // "00000010", the intersection {1}
fmt.Printf("%08b\n", x|y) // "00100110", the union {1, 2, 5}
fmt.Printf("%08b\n", x^y) // "00100100", the symmetric difference {2, 5}
fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
for i := uint(0); i < 8; i++ {
if x&(1<<i) != 0 { // membership test
fmt.Println(i) // "1", "5"
}
}
fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
```

Although Go provides unsigned numbers and arithmetic, we tend to use the signed int form
even for quantities that can’t be negative , such as the length of an array, though uint might
seem a more obv iou s choice. 

``` go
var apples int32 = 1
var oranges int16 = 2
var compote int = apples + oranges // compile error
```

This type mismatch can be fixed in several ways, most directly by converting everything to a
common type:

``` go
var compote = int(apples) + int(oranges)
```

As describedin Section 2.5, for every type T, the conversion operation T(x) converts the value
x to type T if the conversion is allowed. Many integer-to-integer conversions do not entail any
change in value; they just tell the compiler how to interpret a value.

``` go
f := 3.141 // a float64
i := int(f)
fmt.Println(f, i) // "3.141 3"
f=1.99
fmt.Println(int(f)) // "1"
```

Float to integer conversion discards any fractional part, truncating toward zero.

``` go
f := 1e100 // a float64
i := int(f) // result is implementation-dependent
```

When printing numbers using the ```fmt``` package, we can control the radix and format with the
%d, %o, and %x verbs, as shown in this example:

```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
// Output:
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEE
```
Runes are printed with %c, or with %q if quoting is desired:

```go
ascii := 'a'
unicode := 'D'
newline := '\n'
fmt.Printf("%d %[1]c %[1]q\n", ascii) // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
fmt.Printf("%d %[1]q\n", newline) // "10 '\n'"
```


