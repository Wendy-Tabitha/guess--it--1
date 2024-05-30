package mathskills

import (
	"math"
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	input := []float64{10, 30, 20, 50, 40}
	expected := 30.0

	result := Average(input)
	roundOutput := math.Round(result)

	if !reflect.DeepEqual(roundOutput, expected) {
		t.Errorf("Expected %v, but got %v", expected, roundOutput)
	}
}
