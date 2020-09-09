package bench

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/action"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

const (
	tagEndTime = "end"
)

// CalcAction 等待一段时间，然后返回Success
type CalcAction struct {
	action.Action
}

type calcMemo struct {
	behavior.BaseMemo
	sum int
}

// CreateMemo create memo for node
func (node *CalcAction) CreateMemo() behavior.Memory {
	m := &calcMemo{
		BaseMemo: behavior.BaseMemo{},
		sum:      0,
	}
	return m
}

func (node *CalcAction) Initialize(cfg *config.BH3Node) error {
	if err := node.Action.Initialize(cfg); err != nil {
		return err
	}
	return nil
}

// Tick function
func (node *CalcAction) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Info().Msg("tick CalcAction")
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	memo := imemo.(*calcMemo)
	memo.sum = sum
	return behavior.StatusSuccess
}

// Enter enter node
func (node *CalcAction) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	// log.Info().Msg("enter CalcAction")
	memo := imemo.(*calcMemo)
	memo.sum = 0
}
