package main

import "fmt"

func printit(i int){
  fmt.Println("i: ", i)
}

func intSeq() func() int {
  i := 0
  printit(i)
  return func() int {
    i += 1
    printit(i)
    return i
  }
}

func main() {
  nextVal := intSeq()
  fmt.Println(nextVal())
  fmt.Println(nextVal())
  fmt.Println(nextVal())
  fmt.Println(nextVal())
  fmt.Println(nextVal())
}
