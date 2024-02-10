package main

func Sum(numbers []int) int {
	sum := 0

	// ignores index and sets value to "number"
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	return nil
}
