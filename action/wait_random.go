package action

import (
	"math/rand"
	"time"

	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/config"
	// "y3pp/log"
)

const (
	tagMaxDuration = "maxDuration"
	tagMinDuration = "minDuration"
)

// WaitRandom 随机等待一段时间，然后返回Success
type WaitRandom struct {
	Action
	rndDuration int64
	minDuration int64
}

// CreateMemo create memo for node
func (node *WaitRandom) CreateMemo() behavior.Memory {
	m := &waitMemo{
		BaseMemo: behavior.BaseMemo{},
		endTime:  0,
	}
	return m
}

func (node *WaitRandom) Initialize(cfg *config.BH3Node) error {
	if err := node.Action.Initialize(cfg); err != nil {
		return err
	}
	max := cfg.GetInt64(tagMaxDuration)
	node.minDuration = cfg.GetInt64(tagMinDuration)
	node.rndDuration = max - node.minDuration
	return nil
}

// Tick function
func (node *WaitRandom) Tick(bb *behavior.Blackboard, imemo behavior.Memory) behavior.BehaviorStatus {
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
func (node *WaitRandom) Enter(bb *behavior.Blackboard, imemo behavior.Memory) {
	// log.Info().Msg("enter wait")
	now := time.Now().UnixNano() / int64(time.Millisecond)
	memo := imemo.(*waitMemo)
	dur := rand.Int63n(node.rndDuration) + node.minDuration
	memo.endTime = now + dur
}
