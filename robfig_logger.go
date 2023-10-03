package golibcron

import "github.com/golibs-starter/golib/log"

type RobfigLogger struct {
	logger log.Logger
}

func NewRobfigLogger(logger log.Logger) *RobfigLogger {
	return &RobfigLogger{logger: logger}
}

func (r RobfigLogger) Info(msg string, keysAndValues ...interface{}) {
	var args []interface{}
	args = append(args, msg)
	args = append(args, keysAndValues...)
	r.logger.Info(args...)
}

func (r RobfigLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	var args []interface{}
	args = append(args, msg)
	args = append(args, keysAndValues...)
	r.logger.WithError(err).Info(args...)
}
