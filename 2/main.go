package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a int32 = 1000000000;
	fmt.Println(a)
	b:= "Shakib Al Hasan"
	fmt.Println(string(b[3]) == "k")
	fmt.Println(utf8.RuneCountInString("Hello, 世界"))
	fmt.Println(len("Hello, 世界"))
	fmt.Println("Heloo", "World")
}