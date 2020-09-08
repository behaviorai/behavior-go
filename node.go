package behavior

import (
	"github.com/billyplus/behavior/config"
	"github.com/billyplus/behavior/log"
)

// Node 用来表示树的一个节点
type Node interface {
	CreateMemo() Memory
	Initialize(cfg *config.BH3Node) error
	Enter(bb *Blackboard, memo Memory)
	Tick(bb *Blackboard, memo Memory) BehaviorStatus
	// Exit 每次完成节点时触发
	Exit(bb *Blackboard, memo Memory)
	AddChild(n *Wrapper)
	EnableDebug()
}

// type NodeWrapper interface {
// 	Execute(bb *Blackboard) BehaviorStatus
// }

// func wrapNode(node Node) *Wrapper {
// 	n := &Wrapper{
// 		Node: node,
// 	}
// 	return n
// }

var (
	_ Node = &BaseNode{}
)

type BaseNode struct {
	debug bool
}

func (node *BaseNode) CreateMemo() Memory {
	return &BaseMemo{}
}

func (node *BaseNode) Initialize(cfg *config.BH3Node) error {
	return nil
}

func (node *BaseNode) Enter(bb *Blackboard, memo Memory) {
}

func (node *BaseNode) Tick(bb *Blackboard, memo Memory) BehaviorStatus {
	return StatusFailed
}

func (node *BaseNode) Exit(bb *Blackboard, memo Memory) {
}

func (node *BaseNode) AddChild(n *Wrapper) {
}

func (node *BaseNode) EnableDebug() {
	node.debug = true
}

type Wrapper struct {
	Node
	index int
}

const (
	tagRunning = "running"
)

// Execute execute node
func (wrapper *Wrapper) Execute(bb *Blackboard) BehaviorStatus {

	if wrapper.Node == nil {
		return StatusFailed
	}
	memo := bb.GetNodeMemo(wrapper.index)

	// get running state
	st := memo.GetStatus()

	if st != IsRunning {
		wrapper.Node.Enter(bb, memo)
		st = IsRunning
		memo.SaveStatus(st)
	}
	if st == IsRunning {
		status := wrapper.Node.Tick(bb, memo)
		if debug {
			log.Info("tick", zap.String("node", wrapper.Node.))
		}
		if status != StatusRunning {
			wrapper.Node.Exit(bb, memo)
			// save running state
			memo.SaveStatus(IsReady)
		}
		return status
	}
	return StatusFailed
}

// Stop stop node if it is running
func (wrapper *Wrapper) Stop(bb *Blackboard) {
	memo := bb.GetNodeMemo(wrapper.index)
	// get running state
	st := memo.GetStatus()
	if st == IsRunning {
		memo := bb.GetNodeMemo(wrapper.index)
		wrapper.Node.Exit(bb, memo)
		// save running state
		memo.SaveStatus(IsReady)
	}
}
