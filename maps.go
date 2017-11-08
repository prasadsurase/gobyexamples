package main

import "fmt"

func main() {
  var m = make(map[string]int)
  fmt.Println("m:", m)

  m["a"] = 1
  m["b"] = 2
  m["c"] = 3
  fmt.Println("m:", m)

  fmt.Println("len m:", len(m))

  delete(m, "a")
  fmt.Println("m:", m)

  a, b := m["b"]
  fmt.Println("a:", a)
  fmt.Println("b:", b)

  a, b = m["a"]
  fmt.Println("a:", a)
  fmt.Println("b:", b)

  var n = map[string]int{"foo": 1, "bar": 2}
  fmt.Println("n:", n)
}
