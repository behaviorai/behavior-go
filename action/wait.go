package action

import (
	"time"

	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

// Wait 等待一段时间，然后返回Success
type Wait struct {
	Action
	duration int64
}

type waitMemo struct {
	behavior.BaseMemo
	endTime int64
}

// CreateMemo create memo for node
func (node *Wait) CreateMemo() behavior.Memory {
	m := &waitMemo{
		BaseMemo: behavior.BaseMemo{},
		endTime:  0,
	}
	return m
}

func (node *Wait) Initialize(cfg *config.BH3Node) error {
	if err := node.Action.Initialize(cfg); err != nil {
		return err
	}
	node.duration = cfg.GetInt64(tagDuration)
	return nil
}

// Tick function
func (node *Wait) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Info().Msg("tick wait")
	memo := imemo.(*waitMemo)
	end := memo.endTime
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if now >= end {
		return behavior.StatusSuccess
	}
	return behavior.StatusRunning
}

// Enter enter node
func (node *Wait) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	// log.Info().Msg("enter wait")
	now := time.Now().UnixNano() / int64(time.Millisecond)
	memo := imemo.(*waitMemo)
	memo.endTime = now + node.duration
}
