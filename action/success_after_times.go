package action

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

type SuccessAfterTimes struct {
	Action
	times int
}

type satMemo struct {
	behavior.BaseMemo
	runTime int
}

// CreateMemo create memo for node
func (node *SuccessAfterTimes) CreateMemo() behavior.Memory {
	m := &satMemo{
		BaseMemo: behavior.BaseMemo{},
		runTime:  0,
	}
	return m
}

func (node *SuccessAfterTimes) Initialize(cfg *config.BH3Node) error {
	if err := node.Action.Initialize(cfg); err != nil {
		return err
	}
	node.times = cfg.GetInt("times")
	return nil
}

// Tick function
func (node *SuccessAfterTimes) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Logger.Info().Msg("SuccessAfterTimes tick")
	memo := imemo.(*satMemo)
	runned := memo.runTime
	if runned >= 3 {
		return behavior.StatusSuccess
	}
	memo.runTime = runned + 1
	return behavior.StatusRunning
}

// Enter enter node
func (node *SuccessAfterTimes) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	memo := imemo.(*satMemo)
	memo.runTime = 0
}
