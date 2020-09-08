package behavior

var (
	logger Logger
)

type Logger interface {
	Info(msg string)
}

func SetLogger(l Logger) {
	logger = l
}

// func Info(msg string, opt ...interface{}) {
// 	if l == nil {
// 		return
// 	}
// 	l.Info(msg, opt)
// }
