package popcount_test

import (
    "testing"
    "golangbyexamples/ch2/popcount"
)

func BitCount(x uint64) int {
    x = x - ((x >> 1) & 0x5555555555555555)
    x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
    x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
    x = x + (x >> 8)
    x = x + (x >> 16)
    x = x + (x >> 32)
    return int(x & 0x7f)
}


func PopCountByClearing(x uint64) int {
    n := 0
    for x != 0 {
        x = x & (x -1)
        n++
    }
    return n
}
