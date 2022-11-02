package main

import (
    "fmt"
    "os"
)

func main(){

    argLength := len(os.Args[1:])
    fmt.Println("Arg length is %d \n", argLength)

    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a)
    }
}
