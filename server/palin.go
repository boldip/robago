package server


import (
  "fmt"
  "os"
)

func isPalin(s string) bool {
  if len(s) <= 1 {
    return true
  }
  if s[0] != s[len(s) - 1] {
    return false
  }
  return isPalin(s[1 : len(s) - 1])
}


func main() {
  fmt.Println(isPalin(os.Args[1]))
}
