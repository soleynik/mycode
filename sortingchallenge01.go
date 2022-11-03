/*
Expand on the previous program so that the struct Person also tracks an integer named, Height (to be the height in cm of a Person). Implement the sort.Sort interface to utilize a sort by Height. Be sure to display the results when you finish.
*/

package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
    Height int
}

func (p Person) String() string {
    return fmt.Sprintf("%s: %d %d", p.Name, p.Age, p.Height)
}

type ByAge []Person

func (a ByAge) Len() int {
    return len(a)
}
func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

type ByHeight []Person

func (a ByHeight) Len() int {
    return len(a)
}
func (a ByHeight) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByHeight) Less(i, j int) bool {
    return a[i].Height < a[j].Height
}

func main() {

    people := []Person{
        {"Bob", 31, 163},
        {"John", 42, 175},
        {"Michael", 17, 190},
        {"Jenny", 26, 180},
    }

    fmt.Println(people)

    sort.Sort(ByAge(people))
    fmt.Println(people)

    sort.Sort(ByHeight(people))
    fmt.Println(people)

}
