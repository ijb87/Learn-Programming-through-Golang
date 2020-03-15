package main

import (
  "fmt"
)

func main() {
  var num1 int
  var num2 int
  
  fmt.Printf("Enter a number.\n")
  fmt.Scanf("%d\n", &num1)
  fmt.Printf("Enter a second number.\n")
  fmt.Scanf("%d\n", &num2)
  
  var sum = num1 + num2
  fmt.Print("The sum of these two numbers is: ", sum)
}