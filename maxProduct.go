package max_product

func maxProduct(nums []int) int {
	result := 0
	max := 0
	secondMax := 0
	for _, val := range nums {
		if val > max {
			secondMax = max
			max = val
		} else if val > secondMax {
			secondMax = val
		}
	}
	result = (max - 1) * (secondMax - 1)
	return result
}
