package event

type ManualStagedProgress struct{}

type Title struct {
	Default      string
	WhileRunning string
	OnSuccess    string
	OnFail       string
}

type Task struct {
	Title   Title
	Context string
}
