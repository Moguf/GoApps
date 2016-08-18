package main
// ex) go run fetchall.go http://www.google.co.jp http://www.google.com http://yahoo.co.jp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url,ch) // start a goroutine
		// concurrent function exectuion
	}
	for range os.Args[1:]{
		fmt.Println(<-ch) // recevie from channel ch
		// channel: communication mechanism between goroutines.

	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch<- fmt.Sprint(err)
		return
	}

	nbytes,err := io.Copy(ioutil.Discard, resp.Body)
	// ioutil.Discard : discard the body of response 
	// Copy returns byte count

	resp.Body.Close()

	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v",url,err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}
