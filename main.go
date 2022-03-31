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

	// store all read records in memory
	scanner := bufio.NewScanner(file)
	minHeap := make(Heap, 0, k+1)

	for scanner.Scan() {
		// scan line
		line := scanner.Text()
		split := strings.Split(line, ",")
		idStr, scoreStr := split[0], split[1]
		// add record to slice
		id, _ := strconv.Atoi(idStr)
		score, _ := strconv.Atoi(scoreStr)

		// k=5
		heap.Push(&minHeap, Record{id, score})
		if minHeap.Len() > k { // k+1
			heap.Pop(&minHeap)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	sort.Slice(minHeap, func(i, j int) bool {
		return minHeap[i].Score > minHeap[j].Score
	})

	// return first k elements
	if len(minHeap) < k {
		return minHeap
	}
	return minHeap[:k]
}

/*		MIN HEAP
        ✅ parent < child

		parent = (child - 1) / 2
		left   = parent * 2 + 1
		right  = parent * 2 + 2

		0  1  2  3  4
       [1, 2, 3, 5, 4]

            1
           / \
          2   3       // complete tree = insert from left to right on each level
         / \
        5   4        // up
                     // down
*/

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
