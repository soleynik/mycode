// to run
// go build ~/flags03.go
// ~/flags03 --help
// ~/flags03 -word=carrot --truthy=true -numb 8675309
// ~/flags03 --spiderman="Giving up crime fighting to program GoLang" toothbrush mapleleaf toaster

package main

import (
    "flag"
    "fmt"
)

func main(){

    wordPtr := flag.String("word", "Don't panic!", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("truthy", false, "a bool")

    var spiderman string
    flag.StringVar(&spiderman, "spiderman", "Friendly neighborhood Spider-Man", "Catchphrase for Spider-Man")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("truthy:", *boolPtr)
    fmt.Println("spiderman:", spiderman)
    fmt.Println("tail:", flag.Args())
}
