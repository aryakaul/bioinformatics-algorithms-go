package main

import (
	"commonfuncs"
	"os"
    "fmt"
)


func ReverseComplement(genome string) string {
    var revcomp string
    for i := len(genome)-1; i >= 0; i-- {
        curr_nuc := string(genome[i])
        if curr_nuc == "A" {
            revcomp += "T"
        } else if curr_nuc == "C" {
            revcomp += "G"
        } else if curr_nuc == "G" {
            revcomp += "C"
        } else if curr_nuc == "T" {
            revcomp += "A"
        } else {
            fmt.Println("WE FUCKED")
            return "FUCK THE POLICE"
        }
    }
    return revcomp
}


func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    genome := filein[0][0]
    fmt.Println(ReverseComplement(genome))
}
