package main

import (
    "fmt"
    "time"
    "github.com/tj/go-spin"
)

func foo() string {
    return "bar"
}

func compute() {
    s := spin.New()
    for i := 0; i < 100; i++ {
        fmt.Printf("\r%s", s.Next())
        time.Sleep(100 * time.Millisecond)
    }

}

func main() {
    compute()
    fmt.Println(foo())
}
