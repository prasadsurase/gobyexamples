package main

import "fmt"

func factorial(n int) int {
  if n <= 1 {
    return 1
  }
  return n * factorial(n-1)
}

func fibonacci(limit, i, j int) int {
  cnt := 0
  i = j
  j = i + j
  if(cnt > limit){
    return limit
  }
  return fibonacci(limit, i, j)
}

func main() {
  fmt.Println("factorial(5):", factorial(5))
  fmt.Println("factorial(6):", factorial(6))
  fmt.Println("factorial(7):", factorial(7))
}
