package main

import "fmt"

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

func main(){

    astro1 := Astro{"Mickey", 23, "entertain", true}
    astro2 := Astro{"Mouse", 12, "laugh", false}
    astro3 := Astro{"Duck", 35, "swim", true}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)
}

