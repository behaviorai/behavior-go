package composite

import "github.com/billyplus/behavior"

// "y3pp/log"

// SerialSelector 每次都按顺序执行所有节点，只要有一个子节点返回success，那么当前节点也返回success，否则返回failure
type SerialSelector struct {
	Composite
}

// Tick tick Serialselector
func (node *SerialSelector) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	var status behavior.BehaviorStatus

	// execute every child each tick
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		if child == nil {
			continue
		}
		status = child.Execute(bb)
		if status == behavior.StatusSuccess {
			return status
		}
	}
	return behavior.StatusFailed
}

// Exit exit node
func (node *SerialSelector) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	// stop each child
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		child.Stop(bb)
	}
}
