package main
import "fmt"

func century(year int) int {
    var result int = 0
    if year > 100 {
        r := int(year / 100)
        d := int(year % 100)
        if d > 0 {
            result = r + 1
        } else {
            result = r
        }
    } else {
        result = 1
    }
    return result
}

func main() {
    fmt.Println(century(int(1990)))
    fmt.Println(century(int(1705)))
    fmt.Println(century(int(1900)))
    fmt.Println(century(int(1601)))
    fmt.Println(century(int(2000)))
    fmt.Println(century(int(89)))
}

