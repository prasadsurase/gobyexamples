package main

import "fmt"

func add(nums ...int) int {
  sum := 0
  for num := range nums {
    sum += num
  }
  return sum
}

func main(){
  fmt.Println("1 + 2 + 3 + 4 + 5:", add(1,2,3,4,5,))
}
