package decorator

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

const (
	tagMaxLoop  = "maxLoop"
	tagLoopTime = "time"
)

type Decorator struct {
	behavior.BaseNode
	child *behavior.Wrapper
}

// Initialize init node
func (node *Decorator) Initialize(cfg *config.BH3Node) error {
	return node.BaseNode.Initialize(cfg)
}

// AddChild add child
func (node *Decorator) AddChild(child *behavior.Wrapper) {
	node.child = child
}

// GetChild get child of decorator
func (node *Decorator) GetChild() *behavior.Wrapper {
	return node.child
}

// Exit exit from running
func (node *Decorator) Exit(bb *behavior.Blackboard, memo behavior.Memory) {
	node.child.Stop(bb)
	return
}
