package decorator

import "github.com/billyplus/behavior"

// AlwaysFailed 除了running, 不论子节点返回结果是什么，都返回failed,
type AlwaysFailed struct {
	Decorator
}

// Tick function
func (node *AlwaysFailed) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child == nil {
		return behavior.StatusFailed
	}
	status := child.Execute(bb)
	if status == behavior.StatusRunning {
		return behavior.StatusRunning
	}
	return behavior.StatusFailed
}
