package main

import (
  "fmt"
)

func greetSomeone(name string) {
  fmt.Printf("Hello to %s\n", name)
}

func goodbyeSomeone(name string) {
  fmt.Printf("Goodbye to %s\n", name)
}

func main() {
  fmt.Printf("Hello World\n")
  greetSomeone("Peter")
  greetSomeone("Ros")
  fmt.Printf("Goodbye Cruel World\n")
  goodbyeSomeone("Peter")
  goodbyeSomeone("Ros")
}