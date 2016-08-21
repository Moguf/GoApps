## Program Structure
### Name
- Name (OK [a-z,A-Z,_,0-9],). Go distinguish an upper-case letter from an lower-case letter.
 - Naming rule -> "camel case", Go prefer "DoSomething" and "dooSomething" to "do_somthing".
 
### Declarations
Four major kinds of declarations: var, const, type, and func.
* built-in types
|int|int8|int16|int32|int64|int16||
|uint|uint8|uint16|uint32|uint64|uintptr|
|float32|float64|complex32|complex64|||
|bool|byte|rune|string|error||

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

