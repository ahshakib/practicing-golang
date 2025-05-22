package main
import "fmt"
func main() {
	var s string
	fmt.Scan(&s)
	mp := make(map[rune]int)
	for _, c := range s {
		mp[c]++
	}
	if len(mp) == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}