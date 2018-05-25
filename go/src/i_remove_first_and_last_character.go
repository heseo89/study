package main
import "fmt"

func RemoveChar(word string) string {
    fmt.Println(word[1:])
}

func main() {
    fmt.Println(RemoveChar("eloquent"))
}

