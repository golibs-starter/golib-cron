# Golib Cron

Cron Job for Golib project.

### Setup instruction

Base setup, see [GoLib Instruction](https://github.com/golibs-starter/golib#readme)

Both `go get` and `go mod` are supported.

```shell
go get github.com/golibs-starter/golib-cron
```

### Usage

Using `fx.Option` to include dependencies for injection.

```go
package main

import (
    "context"
    golibcron "github.com/golibs-starter/golib-cron"
    "github.com/golibs-starter/golib/log"
    "go.uber.org/fx"
)

type YourFirstJob struct {
}

func NewYourFirstJob() golibcron.Job {
    return &YourFirstJob{}
}

func (y YourFirstJob) Run(ctx context.Context) {
    log.Infoc(ctx, "Job started")
    // By using context, all logs will be printed will job name and run id. Eg:
    // {..."msg":"Job started","job_meta":{"name":"YourFirstJob","run_id":"5fd8bc9e-6250-11ee-b2a1-448a5b97ab48"}}
}

func main() {
    fx.New(
        // When you want to use redis
        golibcron.Opt(),

        // When you want to add new job
        golibcron.ProvideJob(NewYourFirstJob),

        // When you want to enable graceful shutdown.
        golibcron.OnStopHookOpt(),
    )
}

```

### Configuration

```yaml
app:
    # Configuration available for golibcron.Opt()
    cron:
        jobs:
            -   name: YourFirstJob

                #  The Cron Spec pattern requires 5 entries
                #  representing: minute, hour, day of month, month and day of week, in that order.
                #  It accepts
                #   - Standard crontab specs, e.g. "* * * * ?"
                #   - Descriptors, e.g. "@midnight", "@every 1h30m"
                # Check the cron pattern at: https://en.wikipedia.org/wiki/Cron
                spec: "@every 1m"

                # When you want to disable job
                disabled: false

            -   name: YourSecondJob
                spec: "* * * * *"
                disabled: false

            -   name: YourThirdJob
                spec: "* * * * *"
                disabled: true
```
