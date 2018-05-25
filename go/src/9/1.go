package main
import (
    "compress/gzip"
    "fmt"
    "io/ioutil"
    "os"
)


func main() {
    file, err := os.Open("hello.txt.gz")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    r, err := gzip.NewReader(file)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer r.Close()

    b, err := ioutil.ReadAll(r)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(b))

    // file, err := os.OpenFile(
    //     "hello.txt.gz",
    //     os.O_CREATE|os.O_RDWR|os.O_TRUNC,
    //     os.FileMode(0644))
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // defer file.Close()

    // w := gzip.NewWriter(file)

    // defer w.Close()

    // s := "Hello, world!"
    // w.Write([]byte(s))
    // w.Flush()

}

