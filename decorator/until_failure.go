package decorator

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

// UntilFailure 重复执行子节点，直到子节点返回failed，此节点会返回success。或者达到maxLoop次后，此节点返回failed
type UntilFailure struct {
	Decorator
	maxLoop int
}

type ufMemo struct {
	behavior.BaseMemo
	times int
}

// CreateMemo create memo for node
func (node *UntilFailure) CreateMemo() behavior.Memory {
	m := &ufMemo{
		BaseMemo: behavior.BaseMemo{},
		times:    0,
	}
	return m
}

// Initialize init node
func (node *UntilFailure) Initialize(cfg *config.BH3Node) error {
	if err := node.Decorator.Initialize(cfg); err != nil {
		return err
	}
	node.maxLoop = cfg.GetInt(tagLoopTime)
	return nil
}

// Tick function
func (node *UntilFailure) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child == nil {
		return behavior.StatusFailed
	}
	rmemo := memo.(*ufMemo)

	if node.maxLoop <= 0 || rmemo.times < node.maxLoop {
		status := child.Execute(bb)
		// save time
		rmemo.times++

		// return success only when status is failed
		if status == behavior.StatusFailed {
			return behavior.StatusSuccess
		}
		return behavior.StatusRunning
	}
	return behavior.StatusFailed
}

// Enter enter node
func (node *UntilFailure) Enter(bb *behavior.Blackboard, memo behavior.Memory) {
	rmemo := memo.(*ufMemo)
	rmemo.times = 0
}
