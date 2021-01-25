package main

/**
* ./main a bc def
* ./main -s /a bc def
* ./main -n a bc def
* ./main -help
*/

import (
    "fmt"
    "flag"
    "strings"
)


var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    flag.Parse()
    fmt.Printf(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}
