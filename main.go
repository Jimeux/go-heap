package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	hundredK = "./data/records_100k.csv"
	// oku      = "./data/records_oku.csv"
)

func main() {
	filename := hundredK

	fmt.Println("--------getRanking--------")
	ranking := getRanking(5, filename)
	for i, r := range ranking {
		fmt.Printf("[%d] %d %d\n", i+1, r.ID, r.Score)
	}
	fmt.Println("--------------------------")
	fmt.Println("---getRankingOptimized----")

	ranking = getRankingOptimized(5, filename)
	for i, r := range ranking {
		fmt.Printf("[%d] %d %d\n", i+1, r.ID, r.Score)
	}
	fmt.Println("--------------------------")
}

type Record struct {
	ID, Score int
}

// getRanking returns the records with the K highest scores.
func getRanking(k int, filename string) []Record {
	// open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	// store all read records in memory
	records := make([]Record, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// scan line
		line := scanner.Text()
		split := strings.Split(line, ",")
		idStr, scoreStr := split[0], split[1]
		// add record to slice
		id, _ := strconv.Atoi(idStr)
		score, _ := strconv.Atoi(scoreStr)
		records = append(records, Record{id, score})
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	// sort by score
	sort.Slice(records, func(i, j int) bool {
		return records[i].Score > records[j].Score
	})
	// return first k elements
	if len(records) < k {
		return records
	}
	return records[:k]
}

func getRankingOptimized(k int, filename string) []Record {
	// open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// use a min heap of size k+1 because we push elements one-by-one,
	// and Pop (remove the smallest one) each time.
	minHeap := make(Heap, 0, k+1)

	for scanner.Scan() {
		// scan line
		line := scanner.Text()
		split := strings.Split(line, ",")
		idStr, scoreStr := split[0], split[1]
		id, _ := strconv.Atoi(idStr)
		score, _ := strconv.Atoi(scoreStr)

		// push all records to the heap
		heap.Push(&minHeap, Record{id, score})
		if minHeap.Len() > k {
			// when Len reaches k+1, then call Pop (this removes the smallest element)
			heap.Pop(&minHeap)
			// now only the top k elements seen so far remain in the heap
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	// heap order is not guaranteed, so sort
	sort.Slice(minHeap, func(i, j int) bool {
		return minHeap[i].Score > minHeap[j].Score
	})

	// return first k elements
	if len(minHeap) < k {
		return minHeap
	}
	return minHeap[:k]
}

type Heap []Record

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Score < h[j].Score // min heap
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Record))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
