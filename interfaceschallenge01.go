/*
Within the animal interface, create another method, snooze() int, and run your code. It will fail, because tiger and giraffe no longer implement the animal interface. Fix this by adding the new method to both the tiger and giraffe struct to fix this error.
*/

package main

import "fmt"

type animal interface {
    breathe()
    walk()
    snooze()   int    // this must also RETURN int
}

type tiger struct {
     age int
}

func (l tiger) breathe() {
    fmt.Println("tiger breathes")
}

func (l tiger) walk() {
    fmt.Println("tiger walk")
}

func (l tiger) snooze() int {
    fmt.Println("tiger slept for several hours")
    return 10
    }

type giraffe struct {
     age int
}

func (l giraffe) breathe() {
    fmt.Println("giraffe breathes")
}

func (l giraffe) walk() {
    fmt.Println("giraffe walk")
}

func (l giraffe) snooze() int {
    fmt.Println("giraffe slept for a little bit")
    return 5
}

func main() {
    l := tiger{age: 10}
    callBreathe(l)
    callWalk(l)
    callSnooze(l)

    d := giraffe{age: 5}
    callBreathe(d)
    callWalk(d)

}

func callBreathe(a animal) {
    a.breathe()
}

func callWalk(a animal) {
    a.breathe()
}

func callSnooze(a animal) {
    a.snooze()
}
