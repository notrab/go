package main

import (
  "fmt"
  "strings"
)

type Title string

func (t Title) FixCase() string {
  return strings.Title(string(t))
}

func main() {
  name := Title("hello world")
  fixed := name.FixCase()
  fmt.Println(fixed)
}
