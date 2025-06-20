// Problem link: https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/description/?envType=daily-question&envId=2025-06-19

package main
import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
    if len(nums) == 1 {
        return 1
    }
    sort.Ints(nums)
    c, n := 0, len(nums)
    min := nums[0]
    for i := 0; i < n; i++ {
        if nums[i] - min > k {
            c++
            min = nums[i]
        }
        if i == n-1 {
            c++
        }
    }
    return c
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	result := partitionArray(nums, k)
	fmt.Println(result)
}