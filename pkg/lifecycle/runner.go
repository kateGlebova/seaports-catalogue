package lifecycle

type Runnable interface {
	Run()
}

type Runner struct {
	runnables []Runnable
}

func NewRunner(runnables ...Runnable) *Runner {
	return &Runner{runnables: runnables}
}

func (r *Runner) Run() {
	for _, runnable := range r.runnables {
		go runnable.Run()
	}
}
