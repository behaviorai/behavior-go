package composite

import "github.com/billyplus/behavior"

// "y3pp/log"

// ParallelSelector 每次都执行所有节点，只要有一个子节点返回success，那么当前节点也返回success，当所有节点都failure，当前节点也failure
type ParallelSelector struct {
	Composite
}

// // CreateMemo create memo for node
// func (node *ParallelSelector) CreateMemo() behavior.Memory {
// 	m := &seletorMemo{
// 		BaseMemo:    behavior.BaseMemo{},
// 		runningNode: -1,
// 	}
// 	return m
// }

// Tick tick Parallelselector
func (node *ParallelSelector) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	var status behavior.BehaviorStatus
	retStatus := behavior.StatusFailed

	// execute every child each tick
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		if child == nil {
			continue
		}
		status = child.Execute(bb)
		switch status {
		case behavior.StatusRunning:
			if retStatus == behavior.StatusFailed {
				retStatus = behavior.StatusRunning
			}
		case behavior.StatusSuccess:
			retStatus = behavior.StatusSuccess
		}
	}
	return retStatus
}

// Exit exit node
func (node *ParallelSelector) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	// stop each child
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		child.Stop(bb)
	}
}
