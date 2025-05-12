package main
import "fmt"
func main() {
    var t, n int
    fmt.Scan(&t)
    for t > 0 {
        fmt.Scan(&n)
        fmt.Println(n-1)
        t--
    }
}