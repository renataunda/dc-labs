package main

import (
	"fmt"
	"os"
)

func main() {
  name := ""
  
  //foreach, para hacer un for debes nombrar la variable _
  for _, word := range os.Args[1:]{
    name += word + " "
  }
  if len(os.Args) <= 1{
    fmt.Println("Error. Debes poner algun nombre")
  }else{
    fmt.Println("Welcome to the jungle", name)
  }
}
