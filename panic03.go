package main

import "fmt"

func main() {

    magic := make(map[string]string)

    magic["forests"] = "green"
    magic["mountains"] = "red"
    magic["plains"] = "white"

    fmt.Println(magic[100])
}
