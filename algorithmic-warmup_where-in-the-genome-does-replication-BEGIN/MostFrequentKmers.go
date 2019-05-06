package main

import (
	"commonfuncs"
	"os"
    "strconv"
)

func MostFrequentKmers(text string, k int) []string {
    var max_kmers []string
    count := map[string]int{}
    for i := 0; i < len(text)-k+1; i++ {
        curr_kmer := text[i:i+k]
        count[curr_kmer] += 1
    }
    curr_max := 0
    for key, value := range count {
        if value == curr_max {
            max_kmers = append(max_kmers, key)
        }
        if value > curr_max {
            curr_max = value
            max_kmers = nil
            max_kmers = append(max_kmers, key)
        }
    }
    return max_kmers
}

func main() {
    filein, _ := commonfuncs.ReadLines(os.Args[1])
    genome := filein[0][0]
    k, _ := strconv.Atoi(filein[1][0])
    commonfuncs.PrintArrayWithSpaces(MostFrequentKmers(genome, k))
}
