package behavior

// "y3pp/log"

type BehaviorTree struct {
	id       string
	title    string
	root     *Wrapper
	nodelist []Wrapper
}

func (tree *BehaviorTree) Tick(bb *Blackboard) BehaviorStatus {
	// log.Info().Msg("BehaviorTree tick")

	// log.Info().Msgf("root is %+v", tree.root)
	if bb == nil {
		return StatusFailed
	}
	status := tree.root.Execute(bb)
	// log.Logger.Debug().Int("status", int(status)).Msg("BehaviorTree tick")
	return status
}

func (tree *BehaviorTree) NewBlackboard() *Blackboard {
	bb := newBlackboard(len(tree.nodelist))
	for i := 0; i < len(tree.nodelist); i++ {
		bb.nodeMemo[i] = tree.nodelist[i].Node.CreateMemo()
	}
	return bb
}
