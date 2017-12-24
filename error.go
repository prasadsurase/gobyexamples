package main

import (
  "fmt"
  "errors"
)

func sayhi(msg string) (string, error) {
  if(msg == "Hi"){
    return "Greetings human", nil
  }
  return "Idiot", errors.New("Don't you understand English?")
}

func main() {
  msg, err := sayhi("Hi")
  fmt.Println("Hello:", msg)
  fmt.Println("Error:", err)
  msg, err = sayhi("Hola")
  fmt.Println("Hello:", msg)
  fmt.Println("Error:", err)
}
