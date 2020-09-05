package decorator

import "github.com/billyplus/behavior"

// Succeeder 不论子节点返回结果是什么，都返回success
type Succeeder struct {
	Decorator
}

// Tick function
func (node *Succeeder) Tick(bb *behavior.Blackboard, memo behavior.Memory) behavior.BehaviorStatus {
	child := node.GetChild()
	if child != nil {
		child.Execute(bb)
	}
	return behavior.StatusSuccess
}
