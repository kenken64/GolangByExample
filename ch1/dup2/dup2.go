package main

import (
    "fmt"
    "os"
    "bufio"
)

func main(){
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, true)
    }else{
        fmt.Printf("%d\n", len(files))
        for _, arg:= range files{
            f, err:= os.Open(arg)
            if err!=nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, false)
        }
    }
}

func countLines(f*os.File, counts map[string]int , flag bool) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        if flag{
            displayCounts(counts)
        }
    }
    if !flag {
        displayCounts(counts)
    }
}

func displayCounts(counts map[string]int) {
    for line, n:=range counts{
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
