package main

/*
* ./main https://golang.org https://www.google.com https://godoc.org
*
*/
import (
    "os"
    "time"
    "fmt"
    "io"
    "net/http"
    "io/ioutil"
)
func main(){
    start:=time.Now()
    ch :=make(chan string)
    for _, url:= range os.Args[1:]{
        go fetch(url, ch) //start a goroutine
    }

    for range os.Args[1:]{
        fmt.Printf(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err:= http .Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }
    nbytes, err :=io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() //don't leak resources
    if err != nil{
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    secs :=time.Since(start).Seconds()
    ch <-fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}