package mathskills

func Average(num []float64) float64 {
	total := 0.0
	for _, nb := range num {
		total += nb
	}
	return total / float64(len(num))
}
