## Program Structure
### Name
- Name (OK [a-z,A-Z,_,0-9],). Go distinguish an upper-case letter from an lower-case letter.
-- Naming rule -> "camel case", Go prefer "DoSomething" and "doSomething" to "do_somthing".

### Declarations
Four major kinds of declarations: var, const, type, and func.

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
```

* File open




