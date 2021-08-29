package algos

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var tryRobHouse func(house int, houses []int) int
	cache := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = -1
	}
	tryRobHouse = func(house int, houses []int) int {
		if house >= len(houses) {
			return 0
		}
		if cache[house] > -1 {
			return cache[house]
		}
		result := max(tryRobHouse(house+1, houses), tryRobHouse(house+2, houses)+houses[house])
		cache[house] = result
		return result
	}
	return tryRobHouse(0, nums)
}
