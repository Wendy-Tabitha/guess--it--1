package mathskills


func Nextrange(numbers []float64) (int, int) {
        avg := Average(numbers)
        stdDev := Deviation(numbers)

        // Adjust factor as needed based on your data
        factor := 2.0

        lower := int(avg - factor*stdDev)
        upper := int(avg + factor*stdDev)
        return lower, upper
}
