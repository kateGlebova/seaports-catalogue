package lifecycle

import "sync"

type Stoppable interface {
	Stop() error
}

type Stopper struct {
	stoppables []Stoppable
}

func NewStopper(stoppables ...Stoppable) *Stopper {
	return &Stopper{stoppables: stoppables}
}

func (r *Stopper) Stop() error {
	var e = &multiError{
		errChan:   make(chan error, len(r.stoppables)),
		errors:    make([]string, 0, len(r.stoppables)),
		separator: ", ",
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(r.stoppables))
	for _, r := range r.stoppables {
		go func(x Stoppable) {
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
