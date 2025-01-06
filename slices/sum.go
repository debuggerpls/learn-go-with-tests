package slices

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, v := range numbersToSum {
		if len(v) != 0 {
			sums = append(sums, Sum(v[1:]))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
