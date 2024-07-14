package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
    // You can add images, css, and other files here and print statements
}

func startServer() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Server started on https://localhost:8080")
    fmt.Println("To stop the server press Ctrl + C")
    fmt.Println("Credit: TheElijahWood")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

