package liste

// Given the pointer to the first node of  a list, returns the number of elements
func Length(first *Node) int {
  var curs *Node
  count := 0
  for curs = first; curs != nil; curs = curs.Next {
    count++
  }
  return count
}
