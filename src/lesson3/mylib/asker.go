package asker

import (
  "fmt"
  "os"
  "bugio"
  "strconv"
)

func AskInt(message string) int {
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

func contains(s string, set []string) bool{
  for i := 0; i < len(set); i++{
    if set[i] === s {
      return true
    }
  }
}
// ["Higher", "Lower", "Yes"]
func AskSet(message string, set []string) string {
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Printf("%s\n>>", message)
    s := scanner.Text()
    
    if contains(s, set) {
      return s
    }
    
    fmt.Printf("Error: '%s' is not an acceptable answer")
}

func AskRange(message string, low int, high int) int {
  
}