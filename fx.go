package golibcron

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/utils"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func Opt() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewProperties),
		golib.RegisterLogContextExtractor(ContextExtractor),
		fx.Provide(NewDefaultRobfigCron),
		fx.Invoke(RegisterJob),
		fx.Invoke(StartCron),
	)
}

func ProvideJob(jobConstructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{Group: "cron_job", Target: jobConstructor})
}

type RegisterJobIn struct {
	fx.In
	Engine Engine
	Jobs   []Job `group:"cron_job"`
	Props  *Properties
}

func RegisterJob(in RegisterJobIn) error {
	if len(in.Jobs) == 0 {
		log.Infof("No cron jobs found to register")
		return nil
	}
	log.Infof("Registering %d cron jobs", len(in.Jobs))
	registeredCount := 0
	for _, job := range in.Jobs {
		jobName := utils.GetStructShortName(job)
		jobConfig, found := in.Props.GetJob(jobName)
		if !found {
			return fmt.Errorf("spec for job %s not found", jobName)
		}
		if jobConfig.Disabled {
			log.Infof("Not-registered cron job [%s] due by job disabled", jobName)
			continue
		}
		if err := in.Engine.AddJob(jobConfig.Spec, job); err != nil {
			return errors.WithMessagef(err, "cannot register cron job %s with spec: %s", jobName, jobConfig.Spec)
		}
		registeredCount++
		log.Infof("Registered cron job [%s] with spec [%s]", jobName, jobConfig.Spec)
	}
	log.Infof("Registered %d cron jobs", registeredCount)
	return nil
}

func StartCron(c Engine) {
	log.Infof("Cron Engine started")
	c.StartAsync()
}

func OnStopHookOpt() fx.Option {
	return fx.Invoke(OnStopHook)
}

type OnStopCronIn struct {
	fx.In
	Lc   fx.Lifecycle
	Cron Engine
}

func OnStopHook(in OnStopCronIn) {
	in.Lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Infof("Receive stop signal for Cron Engine")
			in.Cron.Stop()
			return nil
		},
	})
}
