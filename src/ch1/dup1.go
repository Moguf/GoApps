package main

import (
	"bufio"    // input, output methods
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// map()  : make hash table.
	// make() : create a new empty map.
	
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}
	for line,n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}
