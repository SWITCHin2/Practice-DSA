// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
type Tree struct{
    val int
    left *Tree
    right *Tree
}

func maxDepth(node *Tree, depth int) int {
    if node == nil {
        return depth
    }
    leftDepth := maxDepth(node.left, depth+1)
    rightDepth := maxDepth(node.right, depth+1)

    return max(leftDepth, rightDepth)
}

func max(a int, b int) int {
    if a >= b{
        return a 
    }else{
        return b
    }
}
func main() {
  fmt.Println("Try programiz.pro")
  
  testTree := &Tree{
    val: 1,
    left: &Tree{
        val: 2,
        left:  &Tree{val: 4},
        right: &Tree{val: 5},
    },
    right: &Tree{
        val: 3,
        left:  &Tree{val: 6},
        right: &Tree{val: 7},
    },
}

  maxDepth := maxDepth(testTree, 0)
  
  fmt.Println(maxDepth)
}