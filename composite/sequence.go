package composite

import "github.com/billyplus/behavior"

// Sequenece 顺序执行节点，只要有一个子节点返回failed或running，当前节点也返回failed或running
type Sequenece struct {
	Composite
}

type seqMemo struct {
	behavior.BaseMemo
	runningNode int
}

// CreateMemo create memo for node
func (node *Sequenece) CreateMemo() behavior.Memory {
	m := &seqMemo{
		BaseMemo:    behavior.BaseMemo{},
		runningNode: -1,
	}
	return m
}

// Tick sequence
func (node *Sequenece) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	memo := imemo.(*seqMemo)
	startIndex := memo.runningNode
	if startIndex < 0 {
		startIndex = 0
	}
	// reset running node
	memo.runningNode = -1

	var status behavior.BehaviorStatus
	for i := startIndex; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		if child == nil {
			continue
		}
		status = child.Execute(bb)

		if status == behavior.StatusRunning {
			// save ruuning state, restart from i next time
			memo.runningNode = i

			return behavior.StatusRunning
		} else if status == behavior.StatusFailed {
			return behavior.StatusFailed
		}
	}
	return behavior.StatusSuccess
}

// Exit exit node
func (node *Sequenece) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*seqMemo)

	startIndex := memo.runningNode
	if startIndex >= 0 {
		child := node.GetChild(startIndex)
		child.Stop(bb)
	}
}
