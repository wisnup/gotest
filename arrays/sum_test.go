package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("for collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected %d but got %d, %v", want, got, numbers)
		}
	})

	t.Run("for collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected %d but got %d, %v", want, got, numbers)
		}
	})
}
