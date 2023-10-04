package golibcron

import (
	"context"
	"github.com/google/uuid"
)

type RobfigJob struct {
	ctx context.Context
	job Job
}

func NewRobfigJob(job Job) *RobfigJob {
	ctx := context.Background()
	jobName := GetJobName(job)
	ctx = context.WithValue(ctx, ContextValueJobName, jobName)
	return &RobfigJob{
		job: job,
		ctx: ctx,
	}
}

func (r RobfigJob) Run() {
	id, _ := uuid.NewUUID()
	ctx := context.WithValue(r.ctx, ContextValueJobRunId, id.String())
	r.job.Run(ctx)
}
