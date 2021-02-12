package main

import (
        "fmt"
        "os"
)

func main() {
        name := ""
        //foreach, para hacer un for debes ponerle nombre a la variable_
        for _, word:=range os.Args[1:]{
                name += word + " "
        }
        if len(os.Args) <= 1{
                fmt.Println("ERROOOOOORRRRRR")
        }else{
                fmt.Println("Welcome to the jungle", name)
        }
}
