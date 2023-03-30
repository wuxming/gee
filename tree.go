package main

//将路由存储到前缀树上，node 为前缀树上的一个节点
type node struct {
	pattern  string  //所注册路由
	part     string  //所注册的路由由 / 分割而成的一部分
	children []*node //子节点
}

//将所注册的路由 pattern 插入的前缀树中
func (n *node) instert(pattern string, parts []string, cur int) {
	//part 全部插入完成，这时候 n 为最后一个节点，为 pattern 赋值
	if len(parts) == cur {
		n.pattern = pattern
		return
	}
	//cur 相当于 parts 中的指针
	part := parts[cur]
	child := n.matchChild(part)
	//如果 child 不为空，则证明有相同的 part 存在，直接插入下一个 part
	//否则就需要新建一个子节点。
	if child == nil {
		child = &node{part: part}
		n.children = append(n.children, child)
	}
	child.instert(pattern, parts, cur+1)
}

//匹配子节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

//根据 http 访问的 path 找到最后的前缀树节点
func (n *node) search(parts []string, cur int) *node {
	if len(parts) == cur {
		return n
	}
	part := parts[cur]
	//获取相匹配的所有子节点
	childs := n.matchChilds(part)
	for _, child := range childs {
		//再依次查询
		res := child.search(parts, cur+1)
		if res != nil {
			return res
		}
	}
	return nil

}

//匹配所有子节点，匹配成功的全部返回
func (n *node) matchChilds(part string) []*node {
	var childs []*node
	for _, child := range n.children {
		if child.part == part {
			childs = append(childs, child)
		}
	}
	return childs
}
