package loop

import (
	"sync"
	"time"

	"github.com/ellemouton/playgrounds/panics/fn"
)

var (
	// loopPanicSubscriber must be passed to all fn.GoWithRecover calls
	// so that an importing package can choose to recover from any panics
	// from within loop.
	loopPanicSubscriber *fn.PanicSubscriber
)

// Main is called when Loop is being started as its own process.
func Main() {
	// In this case, if we want things to behave as they do today, we just
	// don't init the panic subscriber. This will force tha panic recovery
	// library to re-throw the panic.
	loop := New()
	loop.start()
}

type Loop struct {
	wg sync.WaitGroup

	quit chan struct{}
}

func New() *Loop {
	return &Loop{
		quit: make(chan struct{}),
	}
}

// StartAsSubserver is called when Loop is being imported and started within
// another process. This other process may choose to provide a panic subscriber.
func (l *Loop) StartAsSubserver(panicSubscriber *fn.PanicSubscriber) {
	loopPanicSubscriber = panicSubscriber

	l.start()
}

func (l *Loop) Stop() {
	close(l.quit)

	l.wg.Wait()
}

func (l *Loop) start() {
	fn.GoWithRecover(l.loopRoutine1, &l.wg, loopPanicSubscriber)
	fn.GoWithRecover(l.loopRoutine2, &l.wg, loopPanicSubscriber)
}

func (l *Loop) loopRoutine1() {
	time.Sleep(time.Second * 2)

	panic("something something swap")
}

func (l *Loop) loopRoutine2() {
	time.Sleep(time.Second)
}
