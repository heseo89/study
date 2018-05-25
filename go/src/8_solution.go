package main
import "fmt"

func Solution(word string) string {
    var rev_string string

    for i := len(word) -1; i >= 9; i-- {
        rev_string += string(word[i])
    }
    return rev_string
}

func main() {
    fmt.Println(Solution("world"))
    fmt.Println("ss");
}

