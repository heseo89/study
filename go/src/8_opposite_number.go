package main
import "fmt"

func Opposite(value int) int {
    return 0 - value
}

func main() {
    fmt.Println(Opposite(1))
    fmt.Println(Opposite(14))
    fmt.Println(Opposite(-34))
}

