package golibcron

import "context"

type Job interface {
	Run(ctx context.Context)
}

type NamedJob interface {
	Name() string
	Job
}
