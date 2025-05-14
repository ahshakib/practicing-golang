package main
import (
	"fmt"
	"math"
)
var dp [100001]int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func jump(i int, k int, h []int) int {
	if i == 0 {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}
	cost := math.MaxInt32
	
	for j := 1; j <= k; j++ {
		if i - j >= 0 {
			a := jump(i-j, k, h) + abs(h[i] - h[i-j])
			if cost > a {
				cost = a
			}
		}
	}
	
	dp[i] = cost
	return dp[i]
}

func main() {
	for i := range dp {
		dp[i] = -1
	}
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	fmt.Println(jump(n-1, k, h))
}