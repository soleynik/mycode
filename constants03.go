package main

import "fmt"

func main(){
    
    var myFloat32 float32     = 4.5
    var myComplex64 complex64 = 4.5
    fmt.Println(myFloat32)
    fmt.Println(myComplex64)

    type AliasString string

    const myUntypedString     = "Untyped String"
    var uts                   = myUntypedString
    fmt.Println(uts)
}
