package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  select{
  case msg := <-messages:
    fmt.Println("Received:", msg)
  default:
    fmt.Println("None received")
  }

  msg := "hello"
  select{
  case messages <- msg:
    fmt.Println("Sent:", msg)
  default:
    fmt.Println("None sent")
  }

  select{
  case msg := <- messages:
    fmt.Println("Received:", msg)
  case sig := <- signals:
    fmt.Println("Received signal:", sig)
  default:
    fmt.Println("No activity")
  }

  messages <- msg
  var input string
  fmt.Scanln(&input)

}
