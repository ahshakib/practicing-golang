package main
import "fmt"

func main() {
	s := []int{1,2,3,4,5} // slice literal
	fmt.Println(s)
	fmt.Println("Length of slice s:", len(s))
	fmt.Println("Capacity of slice s:", cap(s))
}