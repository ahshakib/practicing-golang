package main
import (
	"fmt"
	"strconv"
)

func toDigits(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func toInt(digits []byte) int {
	num, _ := strconv.Atoi(string(digits))
	return num
}

func swap(arr []byte, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		digits := toDigits(n)
		max := n
		l := len(digits)
		for j := 0; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				swap(digits, j, k)
				for m := 0; m < l; m++ {
					
				}
			}
		}
	}
}