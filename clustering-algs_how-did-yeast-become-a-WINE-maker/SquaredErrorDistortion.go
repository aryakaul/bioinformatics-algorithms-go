package main

import (
    "fmt"
    "math"
    "commonfuncs"
    "os"
)

func EuclidDist (x,y []float64) float64 {
    m := len(x)
    n := len(y)
    if m != n { return -1.0 }
    var d float64 = 0
    for i := 0; i < m; i++ {
        d += math.Pow((x[i]-y[i]), 2)
    }
    return math.Sqrt(d)
}


func MinDistance_datapoint (Datapoint []float64, Centers [][]float64) float64 {
    curr_min := math.MaxFloat64
    for i := 0; i < len(Centers); i++ {
        distance_considered := EuclidDist(Datapoint, Centers[i])
        if curr_min > distance_considered {
            curr_min = distance_considered
        }
    }
    return curr_min
}

func SquaredErrorDistortion (centers [][]float64, data [][]float64) float64 {
    var sum float64
    for i := 0; i < len(data); i++ {
        sum = sum + math.Pow(MinDistance_datapoint(data[i], centers), 2)
    }
    inverse := 1.0/float64(len(data))
    sum = sum * inverse
    return sum
}

func main() {
    filein, err := commonfuncs.ReadLines(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
    if err != nil {
        fmt.Println(err)
    }
    delim := commonfuncs.FindIndexFromArray(filein, "--------")
    centers := filein[1:delim]
    data := filein[delim+1:]
    floatcenters, err := commonfuncs.ConvertStrArrayToFloatArray(centers)
    floatdata, err := commonfuncs.ConvertStrArrayToFloatArray(data)
    fmt.Println(SquaredErrorDistortion(floatcenters, floatdata))
}
