package main

import "fmt"

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

type nasaMission struct {
    people []Astro
    number int
    message string
}

func main(){

    astro1 := Astro{"Mickey", 23, "entertain", true}
    astro2 := Astro{"Mouse", 12, "laugh", false}
    astro3 := Astro{"Duck", 35, "swim", true}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

    p := []Astro{astro1, astro2, astro3}

    fmt.Println(p)

    fmt.Println(p[2].mission)

    s := nasaMission{p, 3, "launch"}

    fmt.Println(s)

    fmt.Printf("%+v", s)
}


