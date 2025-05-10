package main
import (
	"fmt"
	"math"
)
func main() {
	var t, a, b int
	fmt.Scan(&t, &a, &b)
	x := t - (a + b)
	max := int(math.Max(float64(a), float64(b)))
	min := int(math.Min(float64(a), float64(b)))
	if x + min > max {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}