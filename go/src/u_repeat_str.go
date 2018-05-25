package main
import "fmt"

func RepeatStr(repetitions int, value string) string {
    var result string

    for n := 0; n <= repetitions - 1; n ++ {
        for i := 0; i <= len(value) - 1; i++ {
            result += string(value[i])
        }
    }
    return result
}

func main() {
    fmt.Println(RepeatStr(4, "a"))
    fmt.Println(RepeatStr(3, "hello "))
    fmt.Println(RepeatStr(2, "abc"))
}

