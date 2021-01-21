package main

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    //"io"
    "strconv"
    "strings"
)

const (
	TB = 1000000000000
	GB = 1000000000
	MB = 1000000
	KB = 1000
)

func main(){
    fmt.Printf("fetch url\n")
    for _, url :=range os.Args[1:] {
        if !strings.HasPrefix(url, "https"){
            fmt.Printf("Must be Https")
            os.Exit(0)
        }else{
            fmt.Printf("has prefix served as HTTPS \n")
            resp, err := http.Get(url)
            if err != nil {
                fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
                os.Exit(1)
            }
            responseBody := resp.Body
            //nb, err := io.Copy(ioutil.Discard, responseBody)
            b, err := ioutil.ReadAll(responseBody)
            defer resp.Body.Close()
            if err != nil {
                fmt.Fprintf(os.Stderr, "fetch reading: %v", err)
                os.Exit(1)
            }
            //_ = b
            //fmt.Printf("%s\n",lenReadable(int(nb),2))
            if resp.StatusCode == 200 {
                fmt.Printf("%s\n",b)
            }
        }
    }
}


func lenReadable(length int, decimals int) (out string) {
	var unit string
	var i int
	var remainder int

	// Get whole number, and the remainder for decimals
	if length > TB {
		unit = "TB"
		i = length / TB
		remainder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		i = length / GB
		remainder = length - (i * GB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		remainder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		remainder = length - (i * KB)
	} else {
		return strconv.Itoa(length) + " B"
	}

	if decimals == 0 {
		return strconv.Itoa(i) + " " + unit
	}

	// This is to calculate missing leading zeroes
	width := 0
	if remainder > GB {
		width = 12
	} else if remainder > MB {
		width = 9
	} else if remainder > KB {
		width = 6
	} else {
		width = 3
	}

	// Insert missing leading zeroes
	remainderString := strconv.Itoa(remainder)
	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}
	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}

	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}
