package formulas

import (
	"math"
	"testing"
)

func TestCalcAverage(t *testing.T) {
	testcases := []struct {
		name  string
		input []float64
		want  float64
	}{
		{
			name:  "Average Test 1",
			input: []float64{173, 122, 127, 159, 142, 113, 179, 188, 144, 134, 144, 175, 140, 190, 112, 106, 192, 151},
			want:  149.50000000,
		},
		{
			name:  "Average Test 2",
			input: []float64{100, 200, 300, 400, 500},
			want:  300,
		},
		{
			name:  "Average Test 3",
			input: []float64{0, 0, 0, 0, 0},
			want:  0,
		},
		{
			name:  "Average Test 4",
			input: []float64{-1, -2, -3, -4, -5},
			want:  -3,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalcAverage(tc.input)
			if math.Round(got) != math.Round(tc.want) {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestCalcMedian(t *testing.T) {
	testcases := []struct {
		name  string
		input []float64
		want  float64
	}{
		{
			name:  "Median Test 1",
			input: []float64{173, 122, 127, 159, 142, 113, 179, 188, 145, 134, 144, 175, 140, 190, 112, 106, 192, 151},
			want:  144.5,
		},
		{
			name:  "Median Test 2",
			input: []float64{100, 200, 300, 400, 500},
			want:  300,
		},
		{
			name:  "Median Test 3",
			input: []float64{0, 0, 0, 0, 0},
			want:  0,
		},
		{2.000000
			name:  "Median Test 4",
			input: []float64{-1, -2, -3, -4, -5},
			want:  -3,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalcMedian(tc.input)
			if math.Round(got) != math.Round(tc.want) {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestCalcVariance(t *testing.T) {
	testcases := []struct {
		name  string
		input []float64
		want  float64
	}{
		{
			name:  "Variance Test 1",
			input: []float64{173, 122, 127, 159, 142, 113, 179, 188, 145, 134, 144, 175, 140, 190, 112, 106, 192, 151},
			want:  745.802469,
		},
		{
			name:  "Variance Test 2",
			input: []float64{100, 200, 300, 400, 500},
			want:  20000.000000,
		},
		{
			name:  "Variance Test 3",
			input: []float64{0, 0, 0, 0, 0},
			want:  0,
		},
		{
			name:  "Variance Test 4",
			input: []float64{-1, -2, -3, -4, -5},
			want:  2.000000,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalcVariance(tc.input, CalcAverage(tc.input))
			if math.Round(got) != math.Round(tc.want) {
				t.Fatalf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestCalcStdDev(t *testing.T) {
	testcases := []struct {
		name  string
		input []float64
		want  float64
	}{
		{
			name:  "Standard Deviation Test 1",
			input: []float64{173, 122, 127, 159, 142, 113, 179, 188, 144, 134, 144, 175, 140, 190, 112, 106, 192, 151},
			want:  27.319610,
		},
		{
			name:  "Standard Deviation Test 2",
			input: []float64{100, 200, 300, 400, 500},
			want:  141.421356,
		},
		{
			name:  "Standard Deviation Test 3",
			input: []float64{0, 0, 0, 0, 0},
			want:  0,
		},
		{
			name:  "Standard Deviation Test 4",
			input: []float64{-1, -2, -3, -4, -5},
			want:  1.414214,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalcStdDev(CalcVariance(tc.input, CalcAverage(tc.input)))
			if math.Round(got) != math.Round(tc.want) {
				t.Fatalf("got %f, want %f", got, tc.want)
			}
		})
	}
}
