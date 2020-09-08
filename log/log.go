package log

var (
	l Logger
)

type Logger interface {
	Info(msg string, opt ...interface{})
}

func SetLogger(logger Logger) {
	l = logger
}

func Info(msg string, opt ...interface{}) {
	if l == nil {
		return
	}
	l.Info(msg, opt)
}
