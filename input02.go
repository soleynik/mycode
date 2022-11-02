package main

import "fmt"

func main(){

    var name string
    var age int

    fmt.Println("Enter your name & age: ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("%s is %d years old\n", name, age)
}
