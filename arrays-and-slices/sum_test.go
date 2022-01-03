package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	toSum := []int{1, 2, 3}
	sum := Sum(toSum)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d, got %d, given %v", expected, sum, toSum)
	}
}

func checkSums(t testing.TB, actual, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSumAll(t *testing.T) {
	sums := SumAll([]int{1, 2, 3}, []int{9, 38675, 644, 27}, []int{-8, -10, -5, -1})
	expected := []int{6, 39355, -24}

	checkSums(t, sums, expected)
}

func BenchmarkTestAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{9, 38675, 644, 27}, []int{-8, -10, -5, -1})
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("can sum empty slices", func(t *testing.T) {
		emptySliceSum := SumAllTails([]int{})
		expected := []int{0}

		checkSums(t, emptySliceSum, expected)
	})

	t.Run("can sum 3 slices", func(t *testing.T) {
		sums := SumAllTails([]int{1, 2, 3}, []int{9, 38675, 644, 27}, []int{-8, -10, -5, -1})
		expected := []int{5, 39346, -16}

		checkSums(t, sums, expected)
	})
}
