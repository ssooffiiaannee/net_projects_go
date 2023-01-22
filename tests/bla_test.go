package bla 

import (
	// "example.com/projects/test"
	"testing"
	"math"
)

func TestMath(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}


func TestSqrt(t *testing.T) {
	got := math.Sqrt(2)
	if got < 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}