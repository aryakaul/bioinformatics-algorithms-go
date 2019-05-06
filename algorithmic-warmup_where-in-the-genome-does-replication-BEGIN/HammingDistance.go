package main

import (
	"commonfuncs"
	"fmt"
	"os"
)

func HammingDistance(s1, s2 string) int {
    n := len(s1)
    m := len(s2)
    if n != m {
        fmt.Println("Error. Strings are not the same length")
        return -1
    }
    curr_score := 0
    for i := 0; i < n; i++ {
        if s1[i] != s2[i] {
            curr_score += 1
        }
    }
    return curr_score
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    s1 := filein[0][0]
    s2 := filein[1][0]
    fmt.Println(HammingDistance(s1, s2))
}
