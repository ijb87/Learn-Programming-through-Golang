package main

import (
  "fmt"
  "os"
  "bugio"
  "strconv"
)

func askInt(message string) int {
  scanner := bufio.NewScanner(os.Stdin)
  var i int = 1
  for {
    fmt.Printf("%s\n>>", message)
    scanner.Scan()
    s := scanner.Text()
    n, err := strconv.Atoi(s)
    
    if err == nil {
      return n
    }
    fmt.Printf("Error: '%s' is not a number\n", s)
  }
}

func main() {
  var answer int = 35
  
  fmt.Printf("I'm thinking of a number")
  for {
    g := askInt("Take a Guess")
    if g == answer {
      fmt.Println("You Got it")
      return
    }
    if g < answer {
      fmt.Println("Nope: Higher")
    } else {
      fmt.Println("Nope: Lower")
    }
  }
  a := askInt("Choose a number")
  b := askInt("Choose another number")
  c := a + b
  fmt.Printf("%d + %d = %d\n", a, b, c) 
  
  fmt.Printf("You chose : %d", n)
}