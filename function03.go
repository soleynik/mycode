package main

import (
    "fmt"
    "errors"
)

func rollchar(firstName string, lastName string) (string, error){
    if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
        return lastName, errors.New("Vegetables are not suitable last names for heroes.")
    }
    return firstName + " the " + lastName, nil
}

func main(){
    fmt.Println("Welcome to the Character Generator")

    playerChar, err := rollchar("Gandalf", "Turnip")

    if err != nil {
        fmt.Println("Error while spawning your requested character.")
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}
