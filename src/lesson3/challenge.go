package main

import (
  "fmt"
)

func printTo100(){
  var i int = 0
  for {
    i = i + 1
    fmt.Println(i)
    if i == 100 {
      return
    }
  }
}

func main() {
  printTo100()
}