package main

import (
	"commonfuncs"
	"os"
    "strconv"
)

func PatternMatching(genome, pattern string) []string {
    var toret []string
    k := len(pattern)
    for i := 0; i < len(genome)-k+1; i++ {
        if genome[i:i+k] == pattern {
            toret = append(toret, strconv.Itoa(i))
        }
    }
    return toret
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    pattern := filein[0][0]
    genome := filein[1][0]
    commonfuncs.PrintArrayWithSpaces(PatternMatching(genome, pattern))
}
