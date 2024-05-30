package mathskills

import (
	"math"
	"reflect"
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	input := []float64{10, 30, 20, 50, 40}
	expected := 14.0

	result := Deviation(input)
	roundOutput := math.Round(result)

	if !reflect.DeepEqual(roundOutput, expected) {
		t.Errorf("Expected %v, but got %v", expected, roundOutput)
	}
}

