## fundamental Information.
Complie and runs.
```bash
go run helloworld.go

```
Build.
```bash
go build helloworld.go
```
Rewrite souce codes in Go format.

```bash
gofmt helloworld.go
```
### For-loop
```go
// a traditional "while" loop
for condition{
    // ...
}
// a traditional infinite loop
for{
    // ...
}
```
### Initialization
```go
s := ""              // Only used within a function.
var s string         // Defualt initialization.
var s = ""           // Rare case.
var s string = ""    // Redundant.
```
