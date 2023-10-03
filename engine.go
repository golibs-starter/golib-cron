package golibcron

type Engine interface {
	AddJob(spec string, cmd Job) error
	StartSync()
	StartAsync()
	Stop()
}
