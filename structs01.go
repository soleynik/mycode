package main

import "fmt"

type coord struct {
    x int
    y int
}

func main(){

    fmt.Println(coord{1,2})

    zulu := coord{42, 100}
    zulu.y = 180
    fmt.Println(zulu)
    fmt.Println("The x coordinate is:", zulu.x)
    fmt.Println("The y coordinate is:", zulu.y)
}
