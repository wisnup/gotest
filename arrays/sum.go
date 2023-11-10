package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	// length := len(numbers)
	// sums := make([]int, length)

	// for index, arrays := range numbers {
	// 	sums[index] = Sum(arrays)
	// }

	var sums []int
	for _, arrays := range numbers {
		sums = append(sums, Sum(arrays))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {

	var sums []int

	for _, arrays := range numbers {
		if len(arrays) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arrays[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
