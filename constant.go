package behavior

// BehaviorStatus 返回的状态定义
type BehaviorStatus int

// 返回的节点状态
const (
	StatusNone BehaviorStatus = iota
	StatusSuccess
	StatusFailed
	StatusRunning
)

var (
	statusStr = []string{
		StatusNone:    "none",
		StatusSuccess: "success",
		StatusFailed:  "failed",
		StatusRunning: "running",
	}
)

func (s BehaviorStatus) String() string {
	return statusStr[s]
}

// NodeStatus status for node
type NodeStatus int

// node status
const (
	IsReady NodeStatus = iota
	IsRunning
)

// NodeCategory node category
type NodeCategory string

// node category
const (
	CateComposite = "composite"
	CateDecorator = "decorator"
	CateAction    = "action"
	CateCondition = "condition"
)
