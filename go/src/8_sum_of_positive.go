package main
import "fmt"

func PositiveSum(numbers []int) int {
    var total int = 0
    for _, value := range numbers {
        if value > 0 {
            total += value
        }
    }
    return total
}

func main() {
    a := []int{32, 29, 78, 16, 81}
    fmt.Println(PositiveSum(a))

    b := []int{-1, 3, 5, -9, 5}
    fmt.Println(PositiveSum(b))
}

