# Strings

A string is an immutable sequence of bytes. Strings may contain arbitrary data, including 
bytes with value 0, but usually they contain human-readable text. Text strings are conventionally interpreted as UTF-8-encoded sequences of 
Unicode code points (runes).

The built-in ```len``` function returns the number of bytes (not runes) in a string , and the index
operation s[i] retrieves the i-th byte of string s, where 0 ≤ i < len(s).

```go
s := "hello, world"
fmt.Println(len(s)) // "12"
fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
```

The substring operation ```s[i:j]``` yields a new string consisting of the bytes of the original string
starting at index i and continuing up to, but not including, the byte at index j.

Either or both of the i and j operands may be omitted, in which case the default values of 0
(the start of the string) and len(s) (its end) are assumed, respectively

```go
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:]) // "hello, world"
```

The + operator makes a new string by concatenating two strings:
```go
fmt.Println("goodbye" + s[5:]) // "goodbye, world"
```

### 3.5.1. String Literals

A string value can be written as a string literal,a sequence of bytes enclosed in double quotes:

```go
"Hello, BF"
```

Because Go source files are always encoded in UTF-8 and Go text strings are conventionally
interpreted as UTF-8, we can include Unicode code points in string literals.

Within a double-quoted string literal, escape sequences that begin with a backslash \ can be
used to insert arbitrary byte values into the string . One set of escapes handles ASCII control
codes li ke newline, carriage return, and tab

```
\a ‘‘aler t’’ or bell
\b backspace
\f form feed
\n newline
\r carriage return
\t tab
\v vertical tab
\' single quote (only in the rune literal '\'')
\" double quote (only wit hin "..." literals)
\\ backslash
```

### 3.5.2. Unicode

Long ago, life was simple and there was, at least in a parochial view, only one character set to
deal with: ASCII, the American Standard Code for Information Interchange . ASCII, or more
precisely US-ASCII, uses 7 bits to represent 128 ‘‘characters’’: the upper-and lower-case letters
of English, digits, and a variety of punctuation and device-control characters. For much of the
early days of computing , this was a dequate, but it left a very large fraction of the wor ld’s
population unable to use their own writing systems in computers. With the growth of the
Internet, data in myriad languages has become much more common. How can this rich variety be dealt with at all and, if possible, efficiently?

The answer is Unicode (unicode.org), which collects all of the characters in all of the world’s
writing systems, plus accents and other diacritical marks, control codes like tab and carriage
return, and plenty of esoterica, and assigns each one a standard number called a ```Unicode code
point``` or, in Go terminology, ```a rune```.


### 3.5.3. UTF-8

UTF-8 is a variable-length encoding of Unicode code points as bytes. UTF-8 was invented by
Ken Thompson and Rob Pike, two of the creators of Go, and is now a Unicode standard. It
uses between 1 and 4 bytes to represent each rune, but only 1 byte for ASCII characters, and
only 2 or 3 bytes for most runes in common use.

Unicode escapes may also be used in rune literals. These three literals are equivalent:

```
'B' '\u4e16' '\U00004e16'
```

We can test whether one str ing contains another as a prefix:

```go
func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
```
or as a suffix:

```go
func HasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
```

or as a substring:

```go
func Contains(s, substr string) bool {
    for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
            return true
            }
        }
	return false
}
```


### 3.5.4. Strings and Byte Slices

Four standard packages are particularly important for manipulating strings: ```bytes```, ```strings```,
```strconv```, and ```unicode```. The strings package provides many functions for searching, replacing, comparing , trimming, splitting, and joining strings.

The ```basename``` function below was inspired by the Unix shell utility of the same name. In our
version, ```basename(s)``` removes any prefix of s that looks like a file system path with components 
separated by slashes, and it removes any suffix that looks like a file type:

```go
fmt.Println(basename("a/b/c.go")) // "c"
fmt.Println(basename("c.d.go")) // "c.d"
fmt.Println(basename("abc")) // "abc"
```

Strings can be converted to byte slices and back again:

```go
s := "abc"
b := []byte(s)
s2 := string(b)

// Conceptually, the []byte(s) conversion allocates a new byte array holding a copy of the bytes
// of s, and yields a slice that references the entirety of that array.
```

To avoid conversions and unnecessary memoryal location, many of the utility functions in the
bytes package directly parallel their counterparts in the strings package. For example, here
are half a dozen functions from ```strings```:

```go
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```

and the corresponding ones from ```bytes```:

```go
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```


### 3.5.5. Conversions between Strings and Numbers

In addition to conversions between strings, runes, and bytes, it’s often necessary to convert
between numeric values and their string representations. This is done with functions from the
```strconv``` package.

To convert an integer to a string , on e opt ion is to use ```fmt.Sprintf```; 
another is to use the function ```strconv.Itoa(‘‘integer to ASCII’’)```:

```go
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // "123 123"
```

```FormatInt``` and ```FormatUint``` can be used to for mat numbers in a different base:

```go
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
```

To parse a string representing an integer, use the ```strconv``` functions ```Atoi``` or ```ParseInt```, or
```ParseUint``` for unsigned integers:

```go
x, err := strconv.Atoi("123") // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```




