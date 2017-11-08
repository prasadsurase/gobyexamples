package main

import "fmt"

func zeroVal(i int) {
  i = 0
}

func zeroPtr(i *int) {
  *i = 0
}

func main() {
  i := 1
  fmt.Println("i:", i)
  zeroVal(i)
  fmt.Println("i:", i)
  zeroPtr(&i)
  fmt.Println("i:", i)

}
