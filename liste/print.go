package liste

import "fmt"

// Given the pointer to the first node of  a list, prints the list
func Print(first *Node) {
  var curs *Node
  for curs = first; curs != nil; curs = curs.Next {
    if curs.Next == nil {
      fmt.Printf("%s\n", curs.Value)
    } else {
      fmt.Printf("%s -> ", curs.Value)
    }
  }
}
