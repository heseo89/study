package main
import "fmt"


func Multiply(a, b int) int {
    return a * b
}


func main() {
    fmt.Println(Multiply(1, 1))
    fmt.Println(Multiply(1, 0))
    fmt.Println(Multiply(2, 5))
    fmt.Println(Multiply(5, 10))
}

