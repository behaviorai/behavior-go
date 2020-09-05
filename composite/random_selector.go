package composite

import (
	"math/rand"

	"github.com/billyplus/behavior"
)

// RandomSelector 随机顺序执行，随机将children排序，然后执行，其它与RandomSelector相同
type RandomSelector struct {
	Composite
}

type rndSeletorMemo struct {
	behavior.BaseMemo
	runningNode int
	rndIndex    []int
}

// CreateMemo create memo for node
func (node *RandomSelector) CreateMemo() behavior.Memory {
	m := &rndSeletorMemo{
		BaseMemo:    behavior.BaseMemo{},
		runningNode: -1,
		rndIndex:    make([]int, node.CountOfChildren()),
	}
	return m
}

// Tick tick RandomSelector
func (node *RandomSelector) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Logger.Debug().Int("childs", node.CountOfChildren()).Msg("RandomSelector tick")
	memo := imemo.(*rndSeletorMemo)

	startIndex := memo.runningNode
	if startIndex < 0 {
		startIndex = 0
	}
	var status behavior.BehaviorStatus

	memo.runningNode = -1

	for i := startIndex; i < node.CountOfChildren(); i++ {
		idx := memo.rndIndex[i]
		child := node.GetChild(idx)
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

// Enter node
func (node *RandomSelector) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*rndSeletorMemo)
	// 重新随机执行顺序
	for i := len(memo.rndIndex) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		memo.rndIndex[i], memo.rndIndex[j] = memo.rndIndex[j], memo.rndIndex[i]
	}
}

// Exit exit node
func (node *RandomSelector) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*rndSeletorMemo)

	startIndex := memo.runningNode
	if startIndex >= 0 {
		idx := memo.rndIndex[startIndex]
		child := node.GetChild(idx)
		child.Stop(bb)
	}
}
