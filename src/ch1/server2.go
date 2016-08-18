package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var const int

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func hander(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}

func counter(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n",count)
	mu.Unlock()
}
	
