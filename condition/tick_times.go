package condition

import (
	"fmt"

	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

// TickTimes return success when ticked times >= n, else return failure.
type TickTimes struct {
	Condition
	times int
}

type satMemo struct {
	behavior.BaseMemo
	runTime int
}

// CreateMemo create memo for node
func (node *TickTimes) CreateMemo() behavior.Memory {
	m := &satMemo{
		BaseMemo: behavior.BaseMemo{},
		runTime:  0,
	}
	return m
}

func (node *TickTimes) Initialize(cfg *config.BH3Node) error {
	if err := node.Condition.Initialize(cfg); err != nil {
		return err
	}
	node.times = cfg.GetInt("times")
	return nil
}

// Tick function
func (node *TickTimes) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
	// log.Logger.Info().Msg("TickTimes tick")

	memo := imemo.(*satMemo)
	memo.runTime++
	fmt.Println("run time=", memo.runTime)
	if memo.runTime >= node.times {
		memo.runTime = 0
		return behavior.StatusSuccess
	}
	return behavior.StatusFailed
}

// Enter enter node
func (node *TickTimes) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	// memo := imemo.(*satMemo)
	// memo.runTime = 0
}
