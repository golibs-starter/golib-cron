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

type Properties struct {
	Jobs []struct {
		Name string
		Spec string
	}
	jobSpecMap map[string]string
}

func (o *Properties) PostBinding() error {
	o.jobSpecMap = make(map[string]string)
	for _, spec := range o.Jobs {
		if len(o.jobSpecMap[spec.Name]) != 0 {
			return fmt.Errorf("duplicated cron job spec: %s", spec.Name)
		}
		o.jobSpecMap[spec.Name] = spec.Spec
	}
	return nil
}

func (o *Properties) GetSpec(jobName string) (spec string, found bool) {
	spec, found = o.jobSpecMap[jobName]
	return
}

func (o *Properties) Prefix() string {
	return "app.cron"
}
