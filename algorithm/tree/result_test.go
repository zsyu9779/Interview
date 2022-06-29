/*
 * Copyright (C) 2022 Baidu, Inc. All Rights Reserved.
 */
package tree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Right = nil
	root.Left = &TreeNode{
		Val: 2,
	}
	fmt.Println(DiameterOfBinaryTree(root))
}
