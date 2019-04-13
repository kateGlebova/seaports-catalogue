package lifecycle

import "sync"

type Runnable interface {
	Run()
	Stop() error
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

func (r *Runner) Stop() error {
	var e = &multiError{
		errChan:   make(chan error, len(r.runnables)),
		errors:    make([]string, 0, len(r.runnables)),
		separator: ", ",
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(r.runnables))
	for _, r := range r.runnables {
		go func(x Runnable) {
			e.errChan <- x.Stop()
			wg.Done()
		}(r)
	}
	wg.Wait()

	close(e.errChan)
	e.collect()
	if e.ok() {
		return nil
	}

	return e
}
