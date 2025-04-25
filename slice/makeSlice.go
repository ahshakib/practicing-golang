package main
import "fmt"

func main() {
	s := make([]int, 5) // create a slice of length 5
	s[1] = 10
	fmt.Println(s) // [0 0 0 0 0]
	fmt.Println("Length of slice s:", len(s)) // Length of slice s: 5
	fmt.Println("Capacity of slice s:", cap(s)) // Capacity of slice s: 5

	s1 := make([]int, 3, 5) // create a slice of length 3 and capacity 5
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	
	fmt.Println(s1) // [1 2 3]
	fmt.Println("Length of slice s1:", len(s1)) // Length of slice s1: 3
	fmt.Println("Capacity of slice s1:", cap(s1)) // Capacity of slice s1: 5
}