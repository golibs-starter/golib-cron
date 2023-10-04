package golibcron

import (
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/log/field"
)

type RobfigLogger struct {
	logger log.Logger
}

func NewRobfigLogger(logger log.Logger) *RobfigLogger {
	return &RobfigLogger{logger: logger.WithField(
		field.String("module", "cron"),
		field.String("engine", "robfig"),
	)}
}

func (r RobfigLogger) Info(msg string, keysAndValues ...interface{}) {
	r.logger.Debug(r.gatherArgs(msg, keysAndValues)...)
}

func (r RobfigLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	r.logger.WithError(err).Error(r.gatherArgs(msg, keysAndValues)...)
}

func (r RobfigLogger) gatherArgs(msg string, keysAndValues []interface{}) []interface{} {
	var args = []interface{}{msg}
	return append(args, keysAndValues...)
}
