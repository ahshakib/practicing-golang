package main
import "fmt"
func main() {
  var c string
  fmt.Scan(&c)
  
  if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
    fmt.Println("vowel")
  } else {
    fmt.Println("consonant")
  }
}