package question2

func addVals(nums []int) []int {
	for i := 0; i < len(nums); i += 2 {
		end := i + 1
		if end > len(nums)-1 {
			continue
		}
		nums[i] = nums[i] + nums[end]
	}
	return nums
}

func removeVals(nums []int) []int {
	halfLength := len(nums) / 2
	for i := 0; i < halfLength; i++ {
		nums = append(nums[:1+i], nums[2+i:]...)
	}
	return nums
}

func AddValsAndRemoveVals(nums []int) []int {
	nums = addVals(nums)
	nums = removeVals(nums)
	return nums
}
