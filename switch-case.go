package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2

  switch {
  case i < 2:
    fmt.Println("less than 2")
  case i == 2:
    fmt.Println("is 2")
  case i > 2:
    fmt.Println("greater than 2")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend!!!")
  default:
    fmt.Println("FML")
  }

  switch {
  case time.Now().Hour() < 12:
    fmt.Println("A.M")
  case time.Now().Hour() >= 12:
    fmt.Println("P.M")
  }

  whoAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("Boolean")
    case int:
      fmt.Println("integer")
    default:
      fmt.Println("who cares ", t)
    }
  }

  whoAmI(true)
  whoAmI("hello")
  whoAmI(1)
}
