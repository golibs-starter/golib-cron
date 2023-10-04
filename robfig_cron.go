package golibcron

import (
	"github.com/golibs-starter/golib/log"
	"github.com/robfig/cron/v3"
)

func NewDefaultRobfigCron(logger log.Logger) Engine {
	rogLogger := NewRobfigLogger(logger)
	c := cron.New(
		cron.WithLogger(rogLogger),
		cron.WithChain(cron.Recover(rogLogger)),
	)
	return &RobfigCron{c: c}
}

func NewRobfigCron(c *cron.Cron) Engine {
	return &RobfigCron{c: c}
}

type RobfigCron struct {
	c *cron.Cron
}

func (r RobfigCron) AddJob(spec string, cmd Job) error {
	_, err := r.c.AddJob(spec, NewRobfigJob(cmd))
	return err
}

func (r RobfigCron) StartSync() {
	r.c.Run()
}

func (r RobfigCron) StartAsync() {
	r.c.Start()
}

func (r RobfigCron) Stop() {
	r.c.Stop()
}
