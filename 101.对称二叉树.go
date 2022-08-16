package leetcode_record

 type TreeNode struct {
	 Val   int
	 Left  *TreeNode
	 Right *TreeNode
 }

 // 第一次
func isSymmetric(root *TreeNode) bool {
if root==nil{
return true
}
left:=root.Left
right:=root.Right
return helper(left,right)
}

func helper(left,right *TreeNode) bool{
if left==nil && right==nil{
return true
}
if left==nil{
return false
}
if right==nil{
return false
}
if left.Val!=right.Val{
return false
}
return helper(left.Left,right.Right) && helper(left.Right,right.Left)
}

