package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	const maxScore = math.MaxInt
	const recordCount = 100_000

	f, err := os.Create("./ktop/data/records_100k.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for id := 1000; id < recordCount; id++ {
		score := rand.Intn(maxScore)
		rec := strconv.Itoa(id) + "," + strconv.Itoa(score) + "\n"
		if _, err := f.WriteString(rec); err != nil {
			panic(err)
		}
	}
}
