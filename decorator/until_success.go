package decorator

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

// UntilSuccess 重复执行子节点，直到子节点返回success，此节点会返回success。或者达到maxLoop次后，此节点返回failed
type UntilSuccess struct {
	Decorator
	maxLoop int
}

type usMemo struct {
	behavior.BaseMemo
	loopTime int
}

// CreateMemo create memo for node
func (node *UntilSuccess) CreateMemo() behavior.Memory {
	m := &usMemo{
		BaseMemo: behavior.BaseMemo{},
		loopTime: 0,
	}
	return m
}

// Initialize init node
func (node *UntilSuccess) Initialize(cfg *config.BH3Node) error {
	if err := node.Decorator.Initialize(cfg); err != nil {
		return err
	}
	node.maxLoop = cfg.GetInt(tagLoopTime)
	return nil
}

// Tick function
func (node *UntilSuccess) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child == nil {
		return behavior.StatusFailed
	}
	rmemo := memo.(*usMemo)

	if node.maxLoop <= 0 || rmemo.loopTime < node.maxLoop {
		status := child.Execute(bb)
		// save time
		rmemo.loopTime++

		// return success only when status is success
		if status == behavior.StatusSuccess {
			return behavior.StatusSuccess
		}
		return behavior.StatusRunning
	}
	return behavior.StatusFailed
}

// Enter enter node
func (node *UntilSuccess) Enter(bb *behavior.Blackboard, memo behavior.Memory) {
	rmemo := memo.(*usMemo)
	rmemo.loopTime = 0
}
