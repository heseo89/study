package main
import "fmt"

// func sum(a int, b int) int {
//     return a + b
// }
// 
// func main() {
//     r := sum(1, 2)
//     fmt.Println(r)
// }

// func sum(a int, b int) (r int) {
//     r = a + b
//     return
// }
// 
// func main() {
//     r := sum(1, 2)
//     fmt.Println(r)
// }

// func SumAndDiff(a int, b int) (int, int) {
//     return a + b, a - b
// }
// 
// func main() {
//     a := SumAndDiff(6, 2)
//     fmt.Println(a);
//     // sum, diff := SumAndDiff(6, 2)
//     // fmt.Println(sum, diff);
// }

// func factorial(n uint64) uint64 {
//     if n == 0 {
//         return 1
//     }
//     return n * factorial(n-1)
// }
// 
// func main() {
//     fmt.Println(factorial(5))
// }

// func sum(a int, b  int) int {
//     return a + b
// }
// 
// func main() {
//     var hello func(a int, b int) int = sum
//     world := sum
//     
//     fmt.Println(hello(1, 2))
//     fmt.Println(world(1, 2))
// }

// func main() {
//     for i := 0; i < 5; i ++ {
//         defer fmt.Printf("%d", i)
//     }
// }

// func main() {
//     panic("Error !!!")
//     fmt.Println("Hello, World!") //  실행 X
// }


func f() {
    defer func() {
        s := recover()
        fmt.Println(s)
    }()
    panic("Error !!!")
}

func main() {
    f()
    fmt.Println("Hello, World")
}
