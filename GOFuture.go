package GOFuture

import (
	"sync"
	"time"
)

type Future struct {
	wg sync.WaitGroup
}

func (f *Future) RunAsync(fn func(wg *sync.WaitGroup)) {
	wg := &f.wg
	wg.Add(1)
	go fn(wg)
}

func (f *Future) Get() {
	f.wg.Wait()
}

func (f *Future) GetWithTimeOut(timeout time.Duration) {
	var done = make(chan bool)
	go func() {
		f.wg.Wait()
		close(done)
	}()
	timeoutChanel := time.After(timeout)
	select {
	case <-done:
	case <-timeoutChanel:
		return
	}
}
