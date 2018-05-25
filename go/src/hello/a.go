package main
import "fmt"

func main() {
    i := 3

    switch i {
    case 4:
        fmt.Println("4444")
        fallthrough
    case 3:
        fmt.Println("333")
        fallthrough
    case 2:
        fmt.Println("222")
        fallthrough
    case 1:
        fmt.Println("111")
        fallthrough
    case 0:
        fmt.Println("000")
    }
}
