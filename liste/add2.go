package liste


// Given the pointer to the first node of  a list, and a new string,
// adds the latter in the correct position to maintain the list in (lexicographic)
// order. Returns the pointer to the new first element.
func AddInOrder(first *Node, x string) (newFirst *Node) {
  var curs, prev * Node

  newNode := new(Node)
  newNode.Value = x

  if first == nil {
    newFirst = newNode
    return
  }
  // The list is not empty
  // Exit condition: curs == nil || curs.Value > x
  prev = nil
  for curs = first; curs != nil && curs.Value <= x; curs = curs.Next {
    prev = curs
  }
  if prev == nil { // Inserting as first element
    newNode.Next = first   // This is the same as curs, in fact
    newFirst = newNode
    return
  }
  // We put the newNode between prev and curs
  prev.Next = newNode
  newNode.Next = curs
  newFirst = first
  return
}
