package ast

// 处理引用逻辑
// 当文件删除时，将对应的content设为空，版本+1
// 当文件新增时，将对应的content设为新的内容，版本+1
// 重命名时，老的content为空，版本+1，新的content为新的内容，版本+1

type IncludeManager struct {
	ref     map[string][]string        // 每个文件直接引用的文件
	realRef map[string]map[string]bool // 每个文件的间接引用
}

func NewIncludeManager() *IncludeManager {
	return &IncludeManager{
		ref: map[string][]string{},
	}
}

func (i *IncludeManager) addRef(from, to string) {
	if from == to {
		return
	}

	if _, ok := i.ref[from]; !ok {
		i.ref[from] = []string{}
	}
	i.ref[from] = append(i.ref[from], to)
}

func (i *IncludeManager) resetRef(from string, to []string) {
	realTo := []string{}
	for _, v := range to {
		if v == from {
			continue
		}
		realTo = append(realTo, v)
	}
	i.ref[from] = realTo
}

func (i *IncludeManager) recaculateRealRef() []string {
	originRef := i.realRef
	i.realRef = map[string]map[string]bool{}
	for from := range i.ref {
		getIndirectRef(i.ref, i.realRef, from)
	}
	var compareFunc func(map[string]bool, map[string]bool) bool = func(a, b map[string]bool) bool {
		if len(a) != len(b) {
			return false
		}
		for k := range a {
			if _, ok := b[k]; !ok {
				return false
			}
		}
		return true
	}

	// 哪个reffrom的引用有变化
	changedRef := []string{}
	for from, toList := range i.realRef {
		if !compareFunc(toList, originRef[from]) {
			changedRef = append(changedRef, from)
		}
		delete(originRef, from)
	}

	// 哪个reffrom被删除了
	for from := range originRef {
		changedRef = append(changedRef, from)
	}

	return changedRef
}

func getIndirectRef(ref map[string][]string, totalRef map[string]map[string]bool, node string) map[string]bool {
	// 当前节点已经被处理过了
	if _, ok := totalRef[node]; ok {
		return totalRef[node]
	}

	listRefNode := ref[node]
	for _, subNode := range listRefNode {
		if node == subNode {
			continue
		}

		// 先添加当前子节点
		if _, ok := totalRef[node]; !ok {
			totalRef[node] = map[string]bool{}
		}
		totalRef[node][subNode] = true

		// 递归添加子节点的子节点
		subNodeRef := getIndirectRef(ref, totalRef, subNode)
		for k := range subNodeRef {
			if k == node {
				continue
			}
			totalRef[node][k] = true
		}
	}
	return totalRef[node]
}

func (i *IncludeManager) getRef(fileName string) []string {
	return i.ref[fileName]
}

func (i *IncludeManager) getWhoRefMe(fileName string) []string {
	res := []string{}
	for k, v := range i.realRef {
		if _, ok := v[fileName]; ok {
			res = append(res, k)
		}
	}
	return res
}

func (i *IncludeManager) getRealRef(fileName string) map[string]bool {
	return i.realRef[fileName]
}
