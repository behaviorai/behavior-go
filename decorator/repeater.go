package decorator

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
)

// Repeater 重复执行子节点，直到达到maxloop
type Repeater struct {
	Decorator
	maxLoop int
}

type repeaterMemo struct {
	behavior.BaseMemo
	loopTime int
}

// CreateMemo create memo for node
func (node *Repeater) CreateMemo() behavior.Memory {
	m := &repeaterMemo{
		BaseMemo: behavior.BaseMemo{},
		loopTime: 0,
	}
	return m
}

// Initialize init node
func (node *Repeater) Initialize(cfg *config.BH3Node) error {
	if err := node.Decorator.Initialize(cfg); err != nil {
		return err
	}
	node.maxLoop = cfg.GetInt(tagMaxLoop)
	return nil
}

// Tick function
func (node *Repeater) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child == nil {
		return behavior.StatusFailed
	}
	rmemo := memo.(*repeaterMemo)

	if node.maxLoop <= 0 || rmemo.loopTime < node.maxLoop {
		child.Execute(bb)
		// save time
		rmemo.loopTime++

		// always return running
		return behavior.StatusRunning
	}
	return behavior.StatusSuccess
}

// Enter enter func
func (node *Repeater) Enter(bb *behavior.Blackboard, memo behavior.Memory) {
	rmemo := memo.(*repeaterMemo)
	rmemo.loopTime = 0
}
