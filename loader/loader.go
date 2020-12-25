package loader

import (
	"github.com/billyplus/behavior"
	"github.com/billyplus/behavior/action"
	"github.com/billyplus/behavior/composite"
	"github.com/billyplus/behavior/condition"
	"github.com/billyplus/behavior/config"
	"github.com/billyplus/behavior/decorator"
)

func init() {
	// composite
	behavior.Register("Selector", &composite.Selector{})
	behavior.Register("RandomSelector", &composite.RandomSelector{})
	behavior.Register("Sequenece", &composite.Sequenece{})
	behavior.Register("RandomSequence", &composite.RandomSequence{})
	behavior.Register("ParallelSequence", &composite.ParallelSequence{})
	behavior.Register("ParallelSelector", &composite.ParallelSelector{})
	// decorator
	behavior.Register("Repeater", &decorator.Repeater{})
	behavior.Register("Inverter", &decorator.Inverter{})
	behavior.Register("Succeeder", &decorator.Succeeder{})
	behavior.Register("UntilFailure", &decorator.UntilFailure{})
	behavior.Register("UntilSuccess", &decorator.UntilSuccess{})
	behavior.Register("AlwaysFailed", &decorator.AlwaysFailed{})
	behavior.Register("AlwaysSuccess", &decorator.AlwaysSuccess{})
	// action
	behavior.Register("Wait", &action.Wait{})
	behavior.Register("WaitRandom", &action.WaitRandom{})
	behavior.Register("WaitFailure", &action.WaitFailure{})
	// condition
	behavior.Register("TickTimes", &condition.TickTimes{})

}

func LoadBeHaviorTreeFromFile(path string) (*behavior.BehaviorManager, error) {
	bh3cfg, err := config.LoadB3File(path)
	if err != nil {
		return nil, err
	}
	mgr, err := behavior.NewBehaviorManager(&bh3cfg.Data)
	if err != nil {
		return nil, err
	}
	return mgr, nil
}
