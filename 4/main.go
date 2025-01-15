package main
import "fmt"

func main() {
	var arr[4] int32
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	var arr2 = [3] int32 {1, 2, 3}
	var arr3 = [...] int32 {1, 2, 3, 4, 5}
	var arrslice [] int32 = [] int32 {3,10, 11}
	fmt.Println(arr[0:2])
	
	fmt.Println(arr2)
	
	fmt.Println(arr3)
	fmt.Printf("Length: %d, Capacity: %d\n", len(arrslice), cap(arrslice))
	arrslice = append(arrslice, 12)
	fmt.Println(arrslice)
	
	fmt.Printf("Length: %d, Capacity: %d\n", len(arrslice), cap(arrslice))
	var arrslice2 []int32 = []int32 {13, 14,15,90, 86}
	fmt.Println(arrslice2)
	arrslice = append(arrslice, arrslice2...)
	fmt.Println(arrslice)
}