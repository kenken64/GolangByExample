package main

import (
    "fmt"
    "os"
    "strconv"
    
    "golangByexamples/ch2/tempconv"
)


func main() {
    for _ , arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.stderr, "cf: %v\n", err)
            os.exit(1)
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s \n"
            f, tempconv.FToC(f), tempconv.CToF(c))
    }
}
