package golibcron

import (
	"context"
	"github.com/golibs-starter/golib/utils"
)

type Job interface {
	Run(ctx context.Context)
}

type NamedJob interface {
	Name() string
	Job
}

func GetJobName(job Job) string {
	jobName := ""
	if namedJob, ok := job.(NamedJob); ok {
		jobName = namedJob.Name()
	} else {
		jobName = utils.GetStructShortName(job)
	}
	return jobName
}
