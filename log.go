package behavior

import "github.com/billyplus/behavior/log"

func SetLogger(l log.Logger) {
	log.SetLogger(l)
	debug = true
}
