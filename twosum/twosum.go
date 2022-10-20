package twosum

func TwoSum(nums []int, target int) []int {
    memo := make(map[int]int)
    for idx, num := range nums {
        if val, ok := memo[num]; ok {
            return []int{val, idx}
        } else {
            memo[target-num] = idx
        }
    }
    return []int{-1, -1}
}
