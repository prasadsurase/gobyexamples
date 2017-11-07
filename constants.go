package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 999999900000
  const d = 3e15 / n
  fmt.Println(d)

  fmt.Println(float64(d))
  fmt.Println(float32(d))

  fmt.Println(math.Pi)
}
