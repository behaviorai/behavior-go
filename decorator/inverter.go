package decorator

import "github.com/billyplus/behavior"

// Inverter 除了running, 不论子节点返回结果是什么，都反转结果（failed->success，success->failed）
type Inverter struct {
	Decorator
}

// Tick function
func (node *Inverter) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child == nil {
		return behavior.StatusFailed
	}
	status := child.Execute(bb)
	switch status {
	case behavior.StatusFailed:
		return behavior.StatusSuccess
	case behavior.StatusSuccess:
		return behavior.StatusFailed
	default:
		return status
	}
}
