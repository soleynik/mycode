package main

import ( 
    "fmt"
    "os"
    "strings"
)

func main(){

    os.Setenv("Roman", "3")
    fmt.Println("Roman env variable: ", os.Getenv("Roman"))

    fmt.Println("RESEARCH:", os.Getenv("RESEARCH"))

    fmt.Println()

    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}
