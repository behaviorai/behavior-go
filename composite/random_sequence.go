package composite

import (
	"math/rand"

	"github.com/billyplus/behavior"
)

// RandomSequence 随机顺序执行，随机将children排序，然后执行，其它与Sequence相同
type RandomSequence struct {
	Composite
}

type rndSequenceMemo struct {
	behavior.BaseMemo
	runningNode int
	rndIndex    []int
}

// CreateMemo create memo for node
func (node *RandomSequence) CreateMemo() behavior.Memory {
	m := &rndSequenceMemo{
		BaseMemo:    behavior.BaseMemo{},
		runningNode: -1,
		rndIndex:    make([]int, node.CountOfChildren()),
	}
	return m
}

// Tick tick RandomSequence
func (node *RandomSequence) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Logger.Debug().Int("childs", node.CountOfChildren()).Msg("RandomSequence tick")
	memo := imemo.(*rndSequenceMemo)

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
		} else if status == behavior.StatusFailed {
			return behavior.StatusFailed
		}
	}
	return behavior.StatusSuccess
}

// Enter node
func (node *RandomSequence) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*rndSequenceMemo)
	// 重新随机执行顺序
	for i := len(memo.rndIndex) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		memo.rndIndex[i], memo.rndIndex[j] = memo.rndIndex[j], memo.rndIndex[i]
	}
}

// Exit exit node
func (node *RandomSequence) Exit(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*rndSequenceMemo)

	startIndex := memo.runningNode
	if startIndex >= 0 {
		idx := memo.rndIndex[startIndex]
		child := node.GetChild(idx)
		child.Stop(bb)
	}
}
