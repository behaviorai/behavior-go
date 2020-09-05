package composite

import "github.com/billyplus/behavior"

// "y3pp/log"

// Selector 顺序执行节点，只要有一个子节点返回success或running，那么当前节点也返回success或running，
type Selector struct {
	Composite
}

type seletorMemo struct {
	behavior.BaseMemo
	runningNode int
}

// CreateMemo create memo for node
func (node *Selector) CreateMemo() behavior.Memory {
	m := &seletorMemo{
		BaseMemo:    behavior.BaseMemo{},
		runningNode: -1,
	}
	return m
}

// Tick tick selector
func (node *Selector) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Logger.Debug().Int("childs", node.CountOfChildren()).Msg("Selector tick")
	memo := imemo.(*seletorMemo)

	startIndex := memo.runningNode
	if startIndex < 0 {
		startIndex = 0
	}
	var status behavior.BehaviorStatus

	memo.runningNode = -1

	for i := startIndex; i < node.CountOfChildren(); i++ {
		child := node.GetChild(i)
		if child == nil {
			continue
			// return behavior.StatusFailed
		}
		status = child.Execute(bb)
		if status == behavior.StatusRunning {
			// save ruuning state, restart from i next time
			memo.runningNode = i

			return behavior.StatusRunning
		} else if status == behavior.StatusSuccess {
			return behavior.StatusSuccess
		}
	}
	return behavior.StatusFailed
}

// Exit exit node
func (node *Selector) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*seletorMemo)

	startIndex := memo.runningNode
	if startIndex >= 0 {
		child := node.GetChild(startIndex)
		child.Stop(bb)
	}
}
