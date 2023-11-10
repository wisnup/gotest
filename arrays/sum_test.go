package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("for a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected %d but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("for multiple collections of any size", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v but got %v", want, got)
		}

	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v but got %v", want, got)
		}
	}

	t.Run("sum tails of multiple collections of any size", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}
		checkSums(t, want, got)
	})

	t.Run("safey sum empty tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, want, got)

	})
}
