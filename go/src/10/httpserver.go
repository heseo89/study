package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    s := "Hello, World!"

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        html := `
        <html>
        <head>
            <title>Hello</title>
            <script type="text/javascript" src="/assets/hello.js"></script>
            <link href="/assets/hello.css" rel="stylesheet" />
        </head>
        <body>
            <span class="hello">` + s + `</span>
        </body>
        </html>
        `

        res.Header().Set("Content-Type", "text/html")
        res.Write([]byte(html))
    })
    http.HandleFunc("/exit", func(res http.ResponseWriter, req *http.Request) {
        fmt.Println("exit")
        os.Exit(0)

    })

    http.Handle(
        "/assets/",
        http.StripPrefix(
            "/assets/",
            http.FileServer(http.Dir("assets")),
        ),
    )

    fmt.Println(http.ListenAndServe(":8787", nil))
}
