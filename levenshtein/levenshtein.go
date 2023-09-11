package levenshtein

import (
	"math"
)

func DistanceGoroutinges(a string, b string) int {
	resultChannel := make(chan int)
	distanceGoroutines(a, 0, b, 0, resultChannel)
	return <-resultChannel
}

func distanceGoroutines(a string, aIndex int, b string, bIndex int, resultChannel chan int) {
	if len(a) == aIndex {
		resultChannel <- len(b) - bIndex
		return
	}

	if len(b) == bIndex {
		resultChannel <- len(a) - aIndex
		return
	}

	if a[aIndex] == b[bIndex] {
		go distanceGoroutines(a, aIndex+1, b, bIndex+1, resultChannel)
		return
	}

	aTailChannel := make(chan int)
	go distanceGoroutines(a, aIndex, b, bIndex+1, aTailChannel)

	bTailChannel := make(chan int)
	go distanceGoroutines(a, aIndex+1, b, bIndex, bTailChannel)

	tailChannel := make(chan int)
	go distanceGoroutines(a, aIndex+1, b, bIndex+1, tailChannel)

	resultChannel <- 1 + int(math.Min(math.Min(float64(<-aTailChannel), float64(<-bTailChannel)), float64(<-tailChannel)))
}

func DistanceDp(a string, b string) int {
	return distanceDp(a, b)
}

func distanceDp(a string, b string) int {
	distances := make([][]int, len(a))
	for i := range distances {
		distances[i] = make([]int, len(b))
	}

	for i := 1; i < len(a); i++ {
		distances[i][0] = i
	}

	for j := 1; j < len(b); j++ {
		distances[0][j] = j
	}

	for j := 1; j < len(b); j++ {
		for i := 1; i < len(a); i++ {
			substitutionCost := 0
			if a[i] != b[j] {
				substitutionCost = 1
			}

			distances[i][j] = int(math.Min(
				float64(math.Min(float64(distances[i-1][j]+1), float64(distances[i][j-1]+1))),
				float64(distances[i-1][j-1]+substitutionCost),
			))
		}
	}

	return distances[len(a)-1][len(b)-1]
}
