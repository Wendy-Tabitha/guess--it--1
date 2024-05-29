package mathskills

func Median(num []float64) float64 {
	n := len(num)
	//median for if the length of the slice is an even number
	if n%2 == 0 {
		return (num[n/2-1] + num[n/2]) / 2
	}
	//median for if the length of the slice is an odd number
	return num[n/2]
}
