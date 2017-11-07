package main

import "fmt"

func main() {
  if 7 % 2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  if n := 9; n < 0 {
    fmt.Println("number is negative")
  } else if n < 10 {
    fmt.Println("number is single digit")
  } else {
    fmt.Println("number has multiple digits")
  }
}
