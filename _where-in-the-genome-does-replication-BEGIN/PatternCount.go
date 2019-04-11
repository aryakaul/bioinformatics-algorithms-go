package main

import (
	"commonfuncs"
	"fmt"
	"os"
)

func PatternCount(text, pattern string) int {
    n := len(text)
    k := len(pattern)
    if n < k {
        fmt.Println("Error. Genome length is smaller than pattern length")
        return -1
    }
    ctr := 0
    for i := 0; i < n-k+1; i++ {
        if text[i:i+k] == pattern {
            ctr += 1
        }
    }
    return ctr
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    genome := filein[0][0]
    pattern := filein[1][0]
    fmt.Println(PatternCount(genome, pattern))
}
