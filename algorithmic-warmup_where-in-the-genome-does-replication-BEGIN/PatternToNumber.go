package main

import (
	"commonfuncs"
	"fmt"
	"os"
)

func BaseConversion(nucleotide string) int {
    if nucleotide == "A" {
        return 1
    } else if nucleotide == "C" {
        return 2
    } else if nucleotide == "G" {
        return 3
    } else if nucleotide == "T" {
        return 4
    }
    return 0
}


func PatternToNumber(pattern string) int {
    curr := 1
    for idx := 0; idx < len(pattern); idx++ {
        curr *= BaseConversion(string(pattern[idx]))
    }
    return curr-1
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    pattern := filein[0][0]
    fmt.Println(PatternToNumber(pattern))
}
