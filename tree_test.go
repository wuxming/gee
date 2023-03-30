package main

import "testing"

func TestTree(t *testing.T) {
	n := &node{}
	pattern := "/test/tree"
	//插入到前缀树
	n.instert(pattern, []string{"test", "tree"}, 0)
	//查询
	resn := n.search([]string{"test", "tree"}, 0)
	if resn.pattern != pattern {
		t.Errorf("查询到的pattern应该是`%s`，而不是`%s`", pattern, resn.pattern)
	}
	pattern = "/test/tree/v2"
	n.instert(pattern, []string{"test", "tree", "v2"}, 0)
	resn2 := n.search([]string{"test", "tree", "v2"}, 0)
	if resn2.pattern != pattern {
		t.Errorf("查询到的pattern应该是`%s`，而不是`%s`", pattern, resn2.pattern)
	}
}
