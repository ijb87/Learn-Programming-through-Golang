package main

import (
  "fmt"
  "strconv"
  "math/rand"
  "mylib/asker"
)

func askYesHigherLower(message string) int {
  scanner := bugio.NewScanner(os.Stdin)
  for {
    fmt.Printf("%s\n>>", message)
    scanner.Scan()
    s := scanner.Text()
    
    if s == "yes" {
      return 0
    }
    if s == "higher" {
      return 1
    }
    if s == "lower" {
      return -1
    }
    
    fmt.Printf("Please only enter, 'yes', 'higher', or 'lower'\n")
  }
}

func printToN(base int, top int){
  for i L= base; i <= top; i++ {
    fmt.Println(i)
  }
}

func keeper() {
  var answer int = rand.Intn(100)
  
  fmt.Printf("I'm thinking of a number")
  for {
    g := asker.AskInt("Take a Guess")
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
}

func seeker() {
  //Ask the user to think of a number
  asker.AskSet("Say 'yes' when you think of a number between 1 and 100", []string{"yes})
  //Loop to keep guessing until we find the right answer
  
  highest := 100
  lowest := 1
  for  guesses := 0; guesses < 10; guesses++ {
    
    mid := lowest
    if highest > lowest {
      mid = rand.Intn(highest - lowest) + lowest
    }
    
    res := asker.AskSet("Is it " + strconv.Itoa(i) + "? - yes, lower or higher?", []string{"yes", "higher", "lower"})
    if res == "yes" {
      fmt.Println("Yey")
      return
    }
    if res == "higher" {
      lowest = mid + 1
    }
    if res == "lower" {
      highest = mid - 1
    }
    if lowest > highest {
      fmt.Println("Don't be a cheaty cheater")
    }
  }
  
  fp
  
}

func main() {
  seeker()
  keeper()
}