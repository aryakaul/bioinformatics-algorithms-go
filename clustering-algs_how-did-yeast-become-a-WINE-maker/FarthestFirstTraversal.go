package main

import (
	"commonfuncs"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func EuclidDist(x, y []float64) float64 {
	m := len(x)
	n := len(y)
	if m != n {
		return -1.0
	}
	var d float64 = 0
	for i := 0; i < m; i++ {
		d += math.Pow((x[i] - y[i]), 2)
	}
	return math.Sqrt(d)
}

func FarthestFirstTraversal(Data [][]float64, k int) [][]float64 {
	var Centers [][]float64
	Centers = append(Centers, Data[0])
	for len(Centers) < k {
		new_center := MaxDistance(Data[:][:], Centers[:][:])
		Centers = append(Centers, new_center)
	}
	return Centers
}

func MaxDistance(Data [][]float64, Centers [][]float64) []float64 {
	curr_max := 0.0
	var selected_point []float64
	for i := 0; i < len(Data); i++ {
		distance_considered := MinDistance_datapoint(Data[i], Centers)
		if curr_max < distance_considered {
			curr_max = distance_considered
			selected_point = Data[i]
		}
	}
	return selected_point
}

func MinDistance_datapoint(Datapoint []float64, Centers [][]float64) float64 {
	curr_min := math.MaxFloat64
	for i := 0; i < len(Centers); i++ {
		distance_considered := EuclidDist(Datapoint, Centers[i])
		if curr_min > distance_considered {
			curr_min = distance_considered
		}
	}
	return curr_min
}

func outputcenters(centers [][]float64) {
	for i := 0; i < len(centers); i++ {
		for j := 0; j < len(centers[i]); j++ {
			fmt.Printf("%.1f ", centers[i][j])
		}
		fmt.Println()
	}
}

func main() {
	filein, err := commonfuncs.ReadLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	data := filein[1:]
	floatdata, err := commonfuncs.ConvertStrArrayToFloatArray(data)
	k, err := strconv.Atoi(strings.Fields(filein[0][0])[0])
	if err != nil {
		fmt.Println(err)
	}
	found_centers := FarthestFirstTraversal(floatdata[:][:], k)
	outputcenters(found_centers)
	//Test Case
	/*
	   Data := [][]float64 {
	       {0.0,0.0} ,
	       {1.0,1.0} ,
	       {2.0,2.0} ,
	       {3.0,3.0} ,
	       {1.0,2.0} ,
	       {0.0,5.0} ,
	       {5.0,5.0} ,
	   }
	   fmt.Println(FarthestFirstTraversal(Data[:][:], 3))
	*/
}
