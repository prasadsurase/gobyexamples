package main

import "fmt"

func sayhi(msg string){
  for i := 0; i < 5; i++ {
    fmt.Println(msg, ":", i)
  }
}

func main(){
  sayhi("Hello")

  go sayhi("Hola")

  go func(msg string){
    fmt.Println(msg)
  }("going")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}
