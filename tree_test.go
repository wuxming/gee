package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	n := &node{}
	pattern, path := "/test/tree", "/test/tree"
	if !testTree(pattern, path, n) {
		t.Errorf(`"%s"匹配"%s"失败，前缀树有错误`, path, pattern)
	}

	pattern, path = "/test/tree/:v2", "/test/tree/123"
	if !testTree(pattern, path, n) {
		t.Errorf(`"%s"匹配"%s"失败，前缀树有错误`, path, pattern)
	}
	//todo 前缀树有错误待处理
	pattern, path = "/test/tree/*v3", "/test/tree/321/abc"
	if !testTree(pattern, path, n) {
		t.Errorf(`"%s"匹配"%s"失败，前缀树有错误`, path, pattern)
	}
	t.Log("测试完成")

}
func testTree(pattern, path string, n *node) bool {
	//插入到前缀树
	parts := parsePatternAndPath(pattern)
	n.instert(pattern, parts, 0)
	//查询
	searchParts := parsePatternAndPath(path)
	searchNode := n.search(searchParts, 0)
	if searchNode != nil {
		return searchNode.pattern == pattern
	} else {
		return false
	}
}
