package BinarySearchTree

import "testing"

func TestInsert(t *testing.T) {
	bst := &bst{}
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	if bst.root == nil {
		t.Errorf("Insert error: root is nil")
	}

	if bst.root.right.right.data != 3 {
		t.Errorf("Insert error: 最大值不在右子樹")
	}
}

func TestFindRec(t *testing.T) {
	bst := &bst{}
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)

	if exists := bst.Find(3); !exists {
		t.Errorf("Find 錯誤: 找不到")
	}

	if exists := bst.Find(4); exists {
		t.Errorf("Find 錯誤: 找到不存在的值 ")
	}
}
