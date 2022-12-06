package liste

import "fmt"

func main() {
  var first *Node

  first = AddInOrder(first, "paolo")
  first = AddInOrder(first, "malato")
  first = AddInOrder(first, "e'")
  first = AddInOrder(first, "stato")
  first = AddInOrder(first, "signor")
  first = AddInOrder(first, "il")
  first = AddInOrder(first, "boldi")

  fmt.Printf("Lunghezza = %d\n", Length(first))
  Print(first)

}
