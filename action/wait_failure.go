package action

import (
	"time"

	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

// WaitFailure 等待一段时间，然后返回Failure
type WaitFailure struct {
	Action
	duration int64
}

type waitFailureMemo struct {
	behavior.BaseMemo
	endTime int64
}

// CreateMemo create memo for node
func (node *WaitFailure) CreateMemo() behavior.Memory {
	m := &waitFailureMemo{
		BaseMemo: behavior.BaseMemo{},
		endTime:  0,
	}
	return m
}

func (node *WaitFailure) Initialize(cfg *config.BH3Node) error {
	if err := node.Action.Initialize(cfg); err != nil {
		return err
	}
	node.duration = cfg.GetInt64(tagDuration)
	return nil
}

// Tick function
func (node *WaitFailure) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Info().Msg("tick waitFailure")
	memo := imemo.(*waitFailureMemo)
	end := memo.endTime
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if now >= end {
		return behavior.StatusFailed
	}
	return behavior.StatusRunning
}

// Enter enter node
func (node *WaitFailure) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	// log.Info().Msg("enter waitFailure")
	now := time.Now().UnixNano() / int64(time.Millisecond)
	memo := imemo.(*waitFailureMemo)
	memo.endTime = now + node.duration
}
