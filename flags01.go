// (1) returns a pointer to a variable
// func String(name string, value string, usage string) *string

// (2) accepts a pointer to a variable
// func StringVar(p *string, name string, value string, usage string)

package main

import (
    "flag"
    "fmt"
)

// to run
// go run ~/flags01.go -n 2
// go run ~/flags01.go --n=4

func main(){

    num := flag.Int("n", 5, "How angry is the Hulk? (# of iterations)")

    flag.Parse()

    n := *num

    i := 0
    
    for i < n {
        fmt.Println("HULK SMASH!")
        i++;
    }

}
