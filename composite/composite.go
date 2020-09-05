package composite

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

var (
	_ behavior.Node = &Composite{}
)

const (
	tagRunningNode = "runningNode"
)

type Composite struct {
	behavior.BaseNode
	children []*behavior.Wrapper
}

func (node *Composite) AddChild(child *behavior.Wrapper) {
	node.children = append(node.children, child)
}

func (node *Composite) GetChild(index int) *behavior.Wrapper {
	return node.children[index]
}

func (node *Composite) Initialize(cfg *config.BH3Node) error {
	return node.BaseNode.Initialize(cfg)
}

func (node *Composite) CountOfChildren() int {
	return len(node.children)
}
