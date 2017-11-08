package main

import "fmt"

func main() {
  s := make([]string, 3)
  fmt.Println("s:", s)
  fmt.Println("s length:", len(s))

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("s:", s)
  fmt.Println("s length:", len(s))

  s = append(s, "d")
  fmt.Println("s:", s)
  fmt.Println("s length:", len(s))

  s = append(s, "e", "f", "g")
  fmt.Println("s:", s)
  fmt.Println("s length:", len(s))

  k := make([]string, len(s))
  copy(k, s)
  fmt.Println("k:", k)

  fmt.Println("s[0:3]:", s[0:3])
  fmt.Println("s[0:7]:", s[0:7])
  fmt.Println("s[5:]:", s[5:])
  fmt.Println("s[:5]:", s[:5])
}
