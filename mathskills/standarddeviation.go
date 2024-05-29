package mathskills

import "math"

func Deviation(num []float64) float64 {
	nb := Variance(num)
	numb := math.Sqrt(nb)
	return numb
}