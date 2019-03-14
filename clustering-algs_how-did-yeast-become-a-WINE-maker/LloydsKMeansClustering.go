package main

import (
    "fmt"
    "math"
    "commonfuncs"
    "os"
    "strings"
    "strconv"
)

func main() {
    filein,err := commonfuncs.ReadLines(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
}
