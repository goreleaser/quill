package event

type ManualStagedProgress struct {
	N     int
	Err   error
	Stage struct {
		Current string
	}
}

func (ManualStagedProgress) SetCompleted() {}

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
