package buildTree

// 429. N叉树的层序遍历
// 题目：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

// 解法
// 遍历子树，子树中包含p和q，返回root

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	inorderMap := make(map[int]int)
	for idx, v := range inorder {
		inorderMap[v] = idx
	}
	var tree *TreeNode

	return build(preorder, inorder, inorderMap, tree)
}

func build(preorder []int, inorder []int, inorderMap map[int]int, tree *TreeNode) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	pre := preorder[0]
	mid := inorderMap[pre]
	tree.Left = build(preorder[1:mid+1], inorder[0:mid], inorderMap, tree)
	tree.Right = build(preorder[mid+1:len(preorder)+1], inorder[mid+1:len(inorder)+1], inorderMap, tree)
	return tree
}

// times: 1
