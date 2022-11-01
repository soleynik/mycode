package main

import "fmt"

type author struct {
    name string
    branch string
    books int
    salary int
}

func (a author) show(){
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

func main(){
    res := author{
        name: "Roman",
        branch: "AZ",
        books: 14,
        salary: 34000,
    }

    res.show()

}
