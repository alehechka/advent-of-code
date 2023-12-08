package utils

func GreatestCommonDenominator(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDenominator(b, a%b)
}

func LeastCommonMultiple(a, b int) int {
	return (a / GreatestCommonDenominator(a, b)) * b
}

func LeastCommonMultipleMany(nums ...int) (lcm int) {
	if len(nums) == 0 {
		return
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lcm = LeastCommonMultiple(nums[0], nums[1])

	for index := 2; index < len(nums); index++ {
		lcm = LeastCommonMultiple(nums[index], lcm)
	}

	return
}
