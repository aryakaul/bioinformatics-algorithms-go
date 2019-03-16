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

func MinDistance_datapoint(Datapoint []float64, Centers [][]float64) int {
	curr_min := math.MaxFloat64
	var center_chosen_idx int
	for i := 0; i < len(Centers); i++ {
		distance_considered := EuclidDist(Datapoint, Centers[i])
		if curr_min > distance_considered {
			curr_min = distance_considered
			center_chosen_idx = i
		}
	}
	return center_chosen_idx
}

func ClusterToCenter(cluster [][]float64) []float64 {
	var aggregate_point []float64
	aggregate_point = make([]float64, len(cluster[0]))
	for datapoint_idx := 0; datapoint_idx < len(cluster); datapoint_idx++ {
		for value_idx := 0; value_idx < len(cluster[datapoint_idx]); value_idx++ {
			aggregate_point[value_idx] += cluster[datapoint_idx][value_idx]
		}
	}
	//fmt.Println(aggregate_point)
	var center_point []float64
	center_point = make([]float64, len(cluster[0]))
	for i := 0; i < len(aggregate_point); i++ {
		center_point[i] = aggregate_point[i] / float64(len(cluster))
	}
	return center_point
}

func CentersToClusters(centers [][]float64, data [][]float64) [][][]float64 {
	var output_clusters [][][]float64
	output_clusters = make([][][]float64, len(centers))
	for j := range output_clusters {
		output_clusters[j] = make([][]float64, 0)
		for k := range output_clusters[j] {
			output_clusters[j][k] = make([]float64, 0)
		}
	}
	for data_idx := 0; data_idx < len(data); data_idx++ {
		center_chosen_idx := MinDistance_datapoint(data[data_idx], centers)
		output_clusters[center_chosen_idx] = append(output_clusters[center_chosen_idx], data[data_idx])
	}
	return output_clusters
}

func Centers_NotEquivalent(x, y [][]float64) bool {
	if len(y) == 0 {
		return true
	}
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] != y[i][j] {
				return true
			}
		}
	}
	return false
}

func LloydKMeansClustering(data [][]float64, k int) [][]float64 {
	var old_centers, new_centers [][]float64
	old_centers = make([][]float64, k)
	copy(old_centers, data[0:k])
	for Centers_NotEquivalent(old_centers, new_centers) {
		copy(old_centers, new_centers)
		new_centers = new_centers[:0]
		clusters_found := CentersToClusters(old_centers, data)
		for cluster_idx := 0; cluster_idx < len(clusters_found); cluster_idx++ {
			cluster_considered := clusters_found[cluster_idx]
			center_created := ClusterToCenter(cluster_considered)
			new_centers = append(new_centers, center_created)
		}
	}
	return old_centers
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
	k, err := strconv.Atoi(strings.Fields(filein[0][0])[0])
	if err != nil {
		fmt.Println(err)
	}
	data := filein[1:]
	floatdata, err := commonfuncs.ConvertStrArrayToFloatArray(data)
	if err != nil {
		fmt.Println(err)
	}
	output_format(LloydKMeansClustering(floatdata, k))
}
