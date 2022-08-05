package algorithm

type Tree struct {
	ID    int
	Name  string
	Pid   int
	Child []*Tree
}

type Input struct {
	ID   int
	Name string
	Pid  int
}

func GetTree(nums []Input) *Tree {
	tree := &Tree{}
	for i := 0; i < len(nums); i++ {
		addNode(nums[i], tree)
	}
	return tree
}

func addNode(input Input, tree *Tree) {
	if tree == nil {
		return
	}
	for i := 0; i < len(tree.Child); i++ {
		if tree.Child[i].ID == input.Pid {
			tree.Child = append(tree.Child, &Tree{
				ID:    input.ID,
				Name:  input.Name,
				Pid:   input.Pid,
				Child: nil,
			})
		}
	}
}
