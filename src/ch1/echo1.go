package main

import (
	"fmt"
	"os"
)
func main(){
	var s, sep string
	// s and sep are string type.
	// s and sep are automaticaly initialized. (String: "" , Numerical: 0)
	fmt.Println(len(os.Args))
	for i:=1; i< len(os.Args);i++{
		s += sep + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
	// comments
}
