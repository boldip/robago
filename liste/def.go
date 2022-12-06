// Provides a definition of the basic datatype required for lists of string
package liste

type Node struct {
  Value string
  Next *Node        // This is a recursive datatype
}
