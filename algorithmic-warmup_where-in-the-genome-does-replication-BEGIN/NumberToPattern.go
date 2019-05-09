package main

import (
	"commonfuncs"
	"fmt"
    "os"
    "strconv"
)

func NumberToSymbol(index int) string {
    if index == 0 {
        return "A"
    } else if index == 1 {
        return "C"
    } else if index == 2 {
        return "G"
    } else if index == 3 {
        return "T"
    }
    return "Z"
}

func NumberToPattern(idx, k int) string {
    if k == 1 {
        return NumberToSymbol(idx)
    }
    quotient := idx/4
    remainder := idx%4
    symbol := NumberToSymbol(remainder)
    PrefixPattern := NumberToPattern(quotient, k-1)
    return PrefixPattern + symbol
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    idx,_ := strconv.Atoi(filein[0][0])
    k,_ := strconv.Atoi(filein[1][0])
    fmt.Println(NumberToPattern(idx, k))
}
