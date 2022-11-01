package main

import "fmt"

func main(){
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[:]

    slice[1] = 7
    fmt.Println("Modifying Slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)

    array[1] = 2
    fmt.Println("Modifying Underlying Array")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)
}
