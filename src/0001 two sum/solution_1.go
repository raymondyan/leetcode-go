package src

/*
 将 nums 的键值存储在map中，使用 target - value 进行查表，可以省去一次for循环，时间复杂度从 2n -> n
*/
func twoSum(nums []int, target int) []int {
	/*
		初始化一个map => implemented by hash
			1. var map[keyType]valueType
			2. make(map[keyType]valueType, capability)
	*/
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return nil
}
