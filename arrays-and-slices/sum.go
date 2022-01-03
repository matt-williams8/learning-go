package arraysandslices

func Sum(toSum []int) (sum int) {
	for _, number := range toSum {
		sum += number
	}

	return
}

func SumAll(toSum ...[]int) (sums []int) {
	sums = make([]int, len(toSum))

	for i, numbers := range toSum {
		sums[i] = Sum(numbers)
	}

	return
}

func SumAllTails(toSumTails ...[]int) (tailSums []int) {
	tailSums = make([]int, len(toSumTails))

	for i, numbers := range toSumTails {
		var numbersTailSum int

		if len(numbers) > 0 {
			numbersTailSum = Sum(numbers[1:])
		}

		tailSums[i] = numbersTailSum
	}

	return
}
