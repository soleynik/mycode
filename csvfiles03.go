package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Open("origin.csv")

    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(f)
    r.Comma = ';'
    r.Comment = '#'

    records, err := r.ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(records)
}
