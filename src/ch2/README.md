## Program Structure
### Name
- Name (OK [a-z,A-Z,_,0-9],). Go distinguish an upper-case letter from an lower-case letter.
 - Naming rule -> "camel case", Go prefer "DoSomething" and "dooSomething" to "do_somthing".
 
### Declarations
Four major kinds of declarations: var, const, type, and func.
* built-in constans
|true|false|iota|nil|
* built-in types
|int|int8|int16|int32|int64|int16||
|uint|uint8|uint16|uint32|uint64|uintptr|
|float32|float64|complex32|complex64|||
|bool|byte|rune|string|error||
* built-in types
|make|len|cap|new|
|append|copy|close|delete|
|complex|real|imag||
|panic|recover|||


### Variables
```go
var name type = expression   // no expression -> 0, false, ""
var a, b, c = true, 2.1, "three"
var f, err = os.Open(file-name)
```
* Short Variable Declarations
```go
name := expression
i, j := 0, 1
i, j = j, i  // tuple assignment ,swap
// := statement is not necessarily declare all
in, err := os.Open(infile)
out, err := os.Create(outfile)
// Error
f, err := os.Open(infile)
f, err := os.Create(outfile)
```
### Pointer
```go
x := 1
var p *int

var p = f()
func f() *int{
    v := 1
    return &v
}
fmt.Println(f() == f()) // "false", return a distinct value.

func incr(p *int) int{
    *p++
    return *p
}
```
* flag package -- set command-line arguments.
```go
var n = flag.Bool("n", false, "help-message")
// flag name, defualt, help-message
flag.Parse()
flag.Args()
```
### new Function

```go
p := new(int)
fmt.Println(*p)
*p = 2
fmt.Println(*p)
```
new(T) = unnamed variable of type T, default zero, and return address(*T).
```go
func newInt() *int{
     return new(int)
}

func newInt() *int{
     var dummy int
     return &dummy
}
```
### tuple assignment
```go
v, ok = m[key]    // v = value , ok = boolean
v, ok = x.(T)
v, ok = <-ch

_, err = io.Copy(dst, src) // discard byte count

medals := []string{"gold", "silver", "bronze"}
```

### Type Declarations
```go
type name underlying-type
```

* open File
```go
import "os"
f, err := os.Open(name)
if err != nil{
   return err
}
f.Close()
```
