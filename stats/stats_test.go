package stats

import "testing"

func TestMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	want := 3.0
	got := Mean(data)
	if got != want {
		t.Errorf("Mean() = %v; want %v", got, want)
	}
}

func TestMedian(t *testing.T) {
	data := []float64{5, 2, 1, 4, 3}
	want := 3.0
	got := Median(data)
	if got != want {
		t.Errorf("Median() = %v; want %v", got, want)
	}
}
