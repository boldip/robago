package liste


// Given the pointer to the first node of  a list, and a new string,
// adds the latter as the FIRST element of the list. Returns the
// pointer to the new first element.
func AddFront(first *Node, x string) (newFirst *Node) {
  newFirst = new(Node)
  newFirst.Value = x
  newFirst.Next = first
  return
}

// Given the pointer to the first node of  a list, and a new string,
// adds the latter as the LAST element of the list. Returns the
// pointer to the new first element.
func AddLast(first *Node, x string) (newFirst *Node) {
  var curs *Node

  newNode := new(Node)
  newNode.Value = x

  if first == nil {
    newFirst = newNode
    return
  }
  // The list is non-empty
  newFirst = first
  for curs = first; curs.Next != nil ; curs = curs.Next {
  }
  // curs is now pointing to the last node of the list
  curs.Next = newNode
  return
}
