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

func partition_function(datapoint []float64, center []float64, stiffness float64) float64 {
	power := -1 * stiffness * EuclidDist(datapoint, center)
	return math.Exp(power)
}

func CreateHiddenMatrix(data [][]float64, centers [][]float64, stiffness float64) [][]float64 {
	n := len(data)
	var HiddenMatrix [][]float64

	//precompute denominators
	var denominators []float64
	denominators = make([]float64, n)
	for i := 0; i < len(data); i++ {
		var curr_denom float64
		for j := 0; j < len(centers); j++ {
			curr_denom += partition_function(data[i], centers[j], stiffness)
		}
		denominators[i] = curr_denom
	}

	//fill in matrix
	for i := 0; i < len(centers); i++ {
		var row_to_create []float64
		for j := 0; j < len(data); j++ {
			row_to_create = append(row_to_create, (partition_function(data[j], centers[i], stiffness) / denominators[j]))
		}
		HiddenMatrix = append(HiddenMatrix, row_to_create)
	}
	return HiddenMatrix
}

func dot_product(a, b []float64) float64 {
	var toret float64
	for i := 0; i < len(a); i++ {
		toret += a[i] * b[i]
	}
	return toret
}

func ones(a int) []float64 {
	var toret []float64
	toret = make([]float64, a)
	for i := 0; i < a; i++ {
		toret[i] = 1
	}
	return toret
}

func SoftClustersToCenters(data [][]float64, HiddenMatrix [][]float64) [][]float64 {
	m := len(data[0])
	var new_centers [][]float64
	new_centers = make([][]float64, 0)

	for i := 0; i < len(HiddenMatrix); i++ {
		curr_row := HiddenMatrix[i]
		one_vector := ones(len(curr_row))
		denominator := dot_product(curr_row, one_vector)
		var center_created []float64
		for j := 0; j < m; j++ {
			var data_dimension []float64
			data_dimension = make([]float64, 0)
			for _, row := range data {
				data_dimension = append(data_dimension, row[j])
			}
			numerator := dot_product(curr_row, data_dimension)
			value := numerator / denominator
			center_created = append(center_created, value)
		}
		new_centers = append(new_centers, center_created)
	}
	return new_centers
}

func SoftKMeansClustering(init_centers [][]float64, data [][]float64, stiffness float64, num_steps int) [][]float64 {
	hm := CreateHiddenMatrix(data, init_centers, stiffness)
	var updated_centers [][]float64
	for i := 0; i < num_steps; i++ {
		updated_centers = updated_centers[:0]
		updated_centers = SoftClustersToCenters(data, hm)
		hm = CreateHiddenMatrix(data, updated_centers, stiffness)
	}
	return updated_centers
}

func output_format(toret [][]float64) {
	for i := 0; i < len(toret); i++ {
		for j := 0; j < len(toret[i]); j++ {
			fmt.Printf("%.3f ", toret[i][j])
		}
		fmt.Println()
	}
}

func main() {
	filein, err := commonfuncs.ReadLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	k, _ := strconv.Atoi(strings.Fields(filein[0][0])[0])
	beta, _ := strconv.ParseFloat(filein[1][0], 64)
	data := filein[2:]
	floatdata, _ := commonfuncs.ConvertStrArrayToFloatArray(data)
	var init_centers [][]float64
	init_centers = make([][]float64, k)
	copy(init_centers, floatdata[:k])
	num_steps := 100
	output_format(SoftKMeansClustering(init_centers, floatdata, beta, num_steps))
}
