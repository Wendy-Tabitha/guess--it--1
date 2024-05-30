package mathskills

import (
	"math"
	"reflect"
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{10, 30, 20, 50, 40}
	expected := 200.0

	result := Variance(input)
	roundOutput := math.Round(result)

	if !reflect.DeepEqual(roundOutput, expected) {
		t.Errorf("Expected %v, but got %v", expected, roundOutput)
	}
}
