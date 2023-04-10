package computerscience

import (
	"math/rand"
)

/*
linearithmic time n log n
*/
func Collector(n int) int {
	isCollected := make([]bool, n)

	count := 0
	distinct := 0
	for distinct < n {
		r := rand.Intn(n)
		count++
		if !isCollected[r] {
			distinct++
			isCollected[r] = true
		}
	}

	return count
}
