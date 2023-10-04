package golibcron

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

func NewProperties(loader config.Loader) (*Properties, error) {
	props := Properties{}
	err := loader.Bind(&props)
	return &props, err
}

type JobConfig struct {
	Name     string
	Spec     string
	Disabled bool
}

type Properties struct {
	Jobs             []JobConfig
	jobMap           map[string]JobConfig
	EnabledDebugMode bool
}

func (o *Properties) PostBinding() error {
	o.jobMap = make(map[string]JobConfig)
	for _, job := range o.Jobs {
		if _, found := o.jobMap[job.Name]; found {
			return fmt.Errorf("duplicated cron job config: %s", job.Name)
		}
		o.jobMap[job.Name] = job
	}
	return nil
}

func (o *Properties) GetJob(jobName string) (job JobConfig, found bool) {
	job, found = o.jobMap[jobName]
	return
}

func (o *Properties) Prefix() string {
	return "app.cron"
}
