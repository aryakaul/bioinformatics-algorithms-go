package main

import (
	"commonfuncs"
	"fmt"
	"math"
	"os"
	//"strings"
	//"strconv"
)

func FindClosestClusters(clusters [][]int, curr_matrix [][]float64) ([]int, []int) {
	n := len(curr_matrix)
	curr_min := math.MaxFloat64
	var cluster1, cluster2 []int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if curr_matrix[i][j] < curr_min {
				curr_min = curr_matrix[i][j]
				cluster1 = clusters[i]
				cluster2 = clusters[j]
			}
		}
	}
	return cluster1, cluster2
}

func MergeClusters(c1 []int, c2 []int) []int {
	var c3 []int
	//c3 = make([]int, len(c1)+len(c2))
	for i := 0; i < len(c1); i++ {
		c3 = append(c3, c1[i])
	}
	for i := 0; i < len(c2); i++ {
		c3 = append(c3, c2[i])
	}
	return c3
}

func ClustersEquivalent(c1 []int, c2 []int) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}

func RemoveUsedClusters(c1_idx int, c2_idx int, old_matrix [][]float64) [][]float64 {
	var new_matrix [][]float64
	for i := 0; i < len(old_matrix); i++ {
		if i == c1_idx || i == c2_idx {
			continue
		}
		var new_row []float64
		for j := 0; j < len(old_matrix[i]); j++ {
			if j == c1_idx || j == c2_idx {
				continue
			}
			new_row = append(new_row, old_matrix[i][j])
		}
		new_matrix = append(new_matrix, new_row)
	}
	return new_matrix
}

func AddNewCluster(cnew []int, old_matrix [][]float64, init_distance [][]float64, clusters [][]int) [][]float64 {
	var new_matrix [][]float64
	new_matrix = make([][]float64, len(old_matrix))
	for i := 0; i < len(old_matrix); i++ {
		new_matrix[i] = make([]float64, len(old_matrix))
	}
	copy(new_matrix, old_matrix)
	var new_row []float64
	for i := 0; i < len(clusters); i++ {
		dist := CalculateDistance(clusters[i], cnew, init_distance)
		new_row = append(new_row, dist)
	}
	new_matrix = append(new_matrix, new_row)
	for i := 0; i < len(new_matrix)-1; i++ {
		new_matrix[i] = append(new_matrix[i], new_row[i])
	}
	return new_matrix
}

func CalculateDistance(c1 []int, c2 []int, init_dist [][]float64) float64 {
	var score float64
	if ClustersEquivalent(c1, c2) {
		return 0.0
	}
	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			score += init_dist[c1[i]-1][c2[j]-1]
		}
	}
	score /= float64(len(c1) * len(c2))
	return score
}

func UpdateClusters(c1 []int, c2 []int, cnew []int, old_clusters [][]int) (int, int, [][]int) {
	var new_clusters [][]int
	var c1_idx int
	for i := 0; i < len(old_clusters); i++ {
		if ClustersEquivalent(c1, old_clusters[i]) {
			c1_idx = i
			break
		}
	}
	var c2_idx int
	for i := 0; i < len(old_clusters); i++ {
		if ClustersEquivalent(c2, old_clusters[i]) {
			c2_idx = i
			break
		}
	}
	for i := 0; i < len(old_clusters); i++ {
		if i == c1_idx || i == c2_idx {
			continue
		}
		new_clusters = append(new_clusters, old_clusters[i])
	}
	new_clusters = append(new_clusters, cnew)
	return c1_idx, c2_idx, new_clusters
}

func OutputClusters(c []int) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%v ", c[i])
	}
	fmt.Println()
}

func HierarchicalClustering(dist_matrix [][]float64) {
	n := len(dist_matrix)
	var clusters [][]int
	for i := 0; i < n; i++ {
		var cluster_i []int
		cluster_i = append(cluster_i, i+1)
		clusters = append(clusters, cluster_i)
	}
	var curr_matrix [][]float64
	curr_matrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		curr_matrix[i] = make([]float64, n)
	}
	copy(curr_matrix, dist_matrix)
	for len(clusters) > 1 {
		c_i, c_j := FindClosestClusters(clusters, curr_matrix)
		c_new := MergeClusters(c_i, c_j)
		OutputClusters(c_new)
		var c_i_idx, c_j_idx int
		c_i_idx, c_j_idx, clusters = UpdateClusters(c_i, c_j, c_new, clusters)
		curr_matrix = RemoveUsedClusters(c_i_idx, c_j_idx, curr_matrix)
		curr_matrix = AddNewCluster(c_new, curr_matrix, dist_matrix, clusters)
	}
}

func main() {
	filein, _ := commonfuncs.ReadLines(os.Args[1])
	data := filein[1:]
	floatdata, _ := commonfuncs.ConvertStrArrayToFloatArray(data)
	HierarchicalClustering(floatdata)
}
