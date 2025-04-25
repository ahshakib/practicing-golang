package main
import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	s := numbers[1:3] // slice from index 1 to 3 (exclusive)
	s[0] = 100 // modifying the slice modifies the original array
	fmt.Println(s) // [2 3]
	fmt.Println("Length of slice s:", len(s))
	fmt.Println("Capacity of slice s:", cap(s))
	s1 := s[1:2]
	fmt.Println(s1) // [3]
	fmt.Println("Length of slice s1:", len(s1))
	fmt.Println("Capacity of slice s1:", cap(s1))
	s1[0] = 200
	fmt.Println("s:", s)
	fmt.Println("s1:", s1)
	fmt.Println("numbers:", numbers)
}