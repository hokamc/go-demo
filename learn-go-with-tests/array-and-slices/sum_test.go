package array_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		testSum(
			[]int{1, 2, 3, 4, 5},
			15,
			t,
		)
	})

	t.Run("array", func(t *testing.T) {
		testSum(
			[]int{1, 2, 3},
			6,
			t,
		)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("some", func(t *testing.T) {
		testSumAllTails([][]int{{1, 2}, {0, 9}}, []int{2, 9}, t)
	})
	t.Run("empty", func(t *testing.T) {
		testSumAllTails([][]int{{}, {3, 4, 5}}, []int{0, 9}, t)
	})
}

func TestSumAllTails(t *testing.T) {
	sum := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}
}

func testSumAllTails(numbersToSum [][]int, expected []int, t *testing.T) {
	sum := SumAllTails(numbersToSum...)

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}
}

func testSum(numbers []int, expected int, t *testing.T) {
	sum := Sum(numbers)

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
