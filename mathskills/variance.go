package mathskills

import "math"

func Variance(num []float64) float64 {
	sum := 0.0
	n := len(num)
	mean := Average(num)
	for _, nb := range num {
		sum += math.Pow(nb - mean, 2) 
	}
	return sum / float64(n)
}
