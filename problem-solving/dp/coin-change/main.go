// coin change problem using memoization
// This program finds the minimum number of coins needed to make a given amount using any number of coins.
// It uses a recursive approach with memoization to optimize the solution.

package main
import "fmt"

var dp[10005]int
var max = 10000000

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(amount int, coins[]int) int {
	if amount == 0 {
		return 0
	}
	if dp[amount] != -1 {
		return dp[amount]
	}
	minCoins := max
	for _, coin := range coins {
		if amount - coin >= 0 {
			minCoins = min(minCoins, coinChange(amount - coin, coins) + 1)
		}
	}
	dp[amount] = minCoins
	return dp[amount]
}

func main() {
	var n, amount int
	fmt.Scan(&n, &amount)
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&coins[i])
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	result := coinChange(amount, coins)
	if result == max {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}
}