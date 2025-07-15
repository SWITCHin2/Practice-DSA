// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
type Tree struct{
    val int
    left *Tree
    right *Tree
}

func inOrder(node *Tree, ans *[]int) {
    if node == nil{
        return
    }
    if node.left != nil{
        inOrder(node.left, ans)
    }
    *ans = append(*ans, node.val)
    if node.right != nil{
        inOrder(node.right, ans)
    }
    return
}

func postOrder(node *Tree, ansPost *[]int){
    if node == nil {
        return
    }
    if node.left != nil{
        postOrder(node.left, ansPost)
    }
    if node.right != nil{
        postOrder(node.right, ansPost)
    }
    *ansPost = append(*ansPost, node.val)
    return
}
func preOrder(node *Tree, ansPre *[]int){
    if node == nil {
        return
    }
    *ansPre = append(*ansPre, node.val)
    if node.left != nil{
        preOrder(node.left, ansPre)
    }
    if node.right != nil{
        preOrder(node.right, ansPre)
    }
    return
}
func levelOrder(node *Tree, level int, ansLevel *[][]int){
    if node == nil{
        return
    }
    if level == len(*ansLevel) {
        *ansLevel = append(*ansLevel, []int{})
    }
    (*ansLevel)[level] = append((*ansLevel)[level], node.val)
    if node.left != nil {
        levelOrder(node.left, level+1, ansLevel)
    }
    if node.right != nil {
        levelOrder(node.right, level+1, ansLevel)
    }
    return
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
  	ansIn := make([]int, 0)
  	ansPost := make([]int, 0)
  	ansPre := make([]int, 0)
  	ansLevel := make([][]int, 0)
  
  	inOrder(testTree, &ansIn)
  	postOrder(testTree, &ansPost)
  	preOrder(testTree, &ansPre)
  	levelOrder(testTree, 0, &ansLevel)
  
  	fmt.Println(ansIn)
 	fmt.Println(ansPost)
  	fmt.Println(ansPre)
  	fmt.Println(ansLevel)
}