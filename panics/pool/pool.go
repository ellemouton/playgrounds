package pool

import (
	"sync"

	"github.com/ellemouton/playgrounds/panics/fn"
)

var (
	// poolPanicSubscriber must be passed to all fn.GoWithRecover calls
	// so that an importing package can choose to recover from any panics
	// from within loop.
	poolPanicSubscriber *fn.PanicSubscriber
)

// Main is called when Loop is being started as its own process.
func Main() {
	// In this case, if we want things to behave as they do today, we just
	// don't init the panic subscriber. This will force tha panic recovery
	// library to re-throw the panic.
	pool := New()
	pool.start()
}

type Pool struct {
	wg sync.WaitGroup

	panicSignal chan struct{}
	quit        chan struct{}
}

func New() *Pool {
	return &Pool{
		panicSignal: make(chan struct{}),
		quit:        make(chan struct{}),
	}
}

// StartAsSubserver is called when Pool is being imported and started within
// another process. This other process may choose to provide a panic subscriber.
func (l *Pool) StartAsSubserver(panicSubscriber *fn.PanicSubscriber) {
	poolPanicSubscriber = panicSubscriber

	l.start()
}

func (l *Pool) Stop() {
	close(l.quit)

	l.wg.Wait()
}

func (l *Pool) start() {
	fn.GoWithRecover(l.poolRoutine1, &l.wg, poolPanicSubscriber)
}

func (l *Pool) poolRoutine1() {
	select {
	case <-l.panicSignal:
		panic("pool panic")
	case <-l.quit:
		return
	}
}

func (l *Pool) TriggerPanicNow() {
	close(l.panicSignal)
}
