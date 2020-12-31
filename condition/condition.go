package condition

import "github.com/billyplus/behavior"

// Condition 根据条件判断返回success或failed
type Condition struct {
	behavior.BaseNode
}

func (node *Condition) String() string {
	return "Condition"
}
