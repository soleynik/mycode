package main

import (
     "fmt"
     "io"
     "os"
     "log"

     "gopkg.in/yaml.v3"
)

type User struct {
     Name       string
     Occupation string
}

func main() {

    fs, err := os.Open("villains.yaml")

    yfile, err := io.ReadAll(fs)

    if err != nil {
        log.Fatal(err)
    }

    data := make(map[string]User)

    // We deserialize the data into the map or users
    err2 := yaml.Unmarshal(yfile, &data)

    if err2  != nil {
        log.Fatal(err2)
    }

    for k, v := range data {
        fmt.Printf("%s: %s\n", k, v)
    }
}



