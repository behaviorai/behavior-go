package composite

import "github.com/billyplus/behavior"

// "y3pp/log"

// Parallel 每次都执行所有节点，永远返回running
type Parallel struct {
	Composite
}

// // CreateMemo create memo for node
// func (node *Parallel) CreateMemo() behavior.Memory {
// 	m := &seletorMemo{
// 		BaseMemo:    behavior.BaseMemo{},
// 		runningNode: -1,
// 	}
// 	return m
// }

// Tick tick Parallel
func (node *Parallel) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
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
func (node *Parallel) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	// stop each child
	for i := 0; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		child.Stop(bb)
	}
}
