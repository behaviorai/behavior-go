package composite

import "github.com/billyplus/behavior"

// "y3pp/log"

// ParallelSequence 每次都会执行所有节点，只要有一个子节点返回failure，那么当前节点也返回failure，当所有节点都成功，当前节点也返回成功
type ParallelSequence struct {
	Composite
}

// Tick tick ParallelSequence
func (node *ParallelSequence) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	var status behavior.BehaviorStatus
	retStatus := behavior.StatusSuccess

	// execute every child each tick
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		if child == nil {
			continue
		}
		status = child.Execute(bb)
		switch status {
		case behavior.StatusRunning:
			if retStatus == behavior.StatusSuccess {
				retStatus = behavior.StatusRunning
			}
		case behavior.StatusFailed:
			retStatus = behavior.StatusFailed
		}
	}
	return retStatus
}

// Exit exit node
func (node *ParallelSequence) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	// stop each child
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		child.Stop(bb)
	}
}
