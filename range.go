package main

import "fmt"

func main() {
  nums := []int{1, 2, 3}
  sum := 0
  for _, num := range nums {
    sum += num
  }

  fmt.Println("nums:", nums)
  fmt.Println("Sum:", sum)

  hash := map[string]int{"foo": 1, "bar": 2, "blah": 3}
  sum = 0

  for _, v := range hash {
    sum += v
  }
  fmt.Println("hash:", hash)
  fmt.Println("Sum:", sum)

  for i, c := range "hello" {
    fmt.Println("i:", i, " c:", c)
  }
}
