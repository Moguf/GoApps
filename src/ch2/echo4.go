package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
// flag name, defualt, help-message
var sep = flag.String("s", " ", "separator")
// flag name, defualt, help-message

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
		
