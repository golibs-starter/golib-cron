package golibcron

import (
	"context"
	"github.com/golibs-starter/golib/log/field"
)

const ContextValueJobName = "job_attributes_name"
const ContextValueJobRunId = "job_attributes_run_id"
const ContextLogJobMeta = "job_meta"

type JobAttributes struct {
	Name  string `json:"name,omitempty"`
	RunId string `json:"run_id,omitempty"`
}

func (c JobAttributes) MarshalLogObject(encoder field.ObjectEncoder) error {
	if c.Name != "" {
		encoder.AddString("name", c.Name)
	}
	if c.RunId != "" {
		encoder.AddString("run_id", c.RunId)
	}
	return nil
}

func ContextExtractor(ctx context.Context) []field.Field {
	jobNameAny := ctx.Value(ContextValueJobName)
	if jobNameAny == nil {
		return nil
	}
	attrs := JobAttributes{}
	if jobName, ok := jobNameAny.(string); ok {
		attrs.Name = jobName
	}
	if attrs.Name == "" {
		return nil
	}
	jobRunIdAny := ctx.Value(ContextValueJobRunId)
	if jobRunIdAny != nil {
		if jobRunId, ok := jobRunIdAny.(string); ok {
			attrs.RunId = jobRunId
		}
	}
	return []field.Field{
		field.Object(ContextLogJobMeta, &attrs),
	}
}
