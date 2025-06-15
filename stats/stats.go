package stats

import (
	"sort"
)

// Mean calculates the average of a slice of floats.
func Mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// Median calculates the median value of a slice of floats.
func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0 // or panic, depending on your API philosophy
	}
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)

	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}
