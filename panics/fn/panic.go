package fn

import (
	"fmt"
	"runtime"
	"sync"
)

// GoWithRecover spins off the given function in a goroutine but first wraps it
// with panic recovery. If a nil PanicSubscriber is provided then any caught
// panic will be re-thrown. An optional waiting group may be passed if a delta
// should be added before starting the goroutine and marked as done on
// completion.
func GoWithRecover(fn func(), wg *sync.WaitGroup, sub *PanicSubscriber) {
	if wg != nil {
		wg.Add(1)
	}

	go func() {
		if wg != nil {
			defer wg.Done()
		}

		defer func() {
			err := recover()
			if err == nil {
				return
			}

			caller := "unknown"
			_, file, no, ok := runtime.Caller(1)
			if ok {
				caller = fmt.Sprintf("%s#%d", file, no)
			}

			log := fmt.Sprintf("panic from: %s: %v", caller, err)

			if sub == nil {
				panic(fmt.Sprint("re-throwing unhandled "+
					"panic: ", log))
			}

			sub.mu.Lock()
			sub.panicLogs = append(sub.panicLogs, log)
			sub.mu.Unlock()

			select {
			// The panicSignal channel has already been closed.
			case <-sub.signal:
				return
			default:
				close(sub.signal)
			}
		}()

		fn()
	}()
}

// PanicSubscriber represents a subscriber to a goroutine panic.
type PanicSubscriber struct {
	mu        sync.Mutex
	panicLogs []string

	signal chan struct{}
}

// NewPanicSubscriber constructs a new PanicSubscribe which may be passed in to
// GoWithRecover to subscribe to a panic from a goroutine.
func NewPanicSubscriber() *PanicSubscriber {
	return &PanicSubscriber{
		signal: make(chan struct{}),
	}
}

// PanicLogs returns all the panics that have been reported to the given
// PanicSubscriber up until this point in time.
func (s *PanicSubscriber) PanicLogs() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.panicLogs
}

// PanicSignal returns a channel that can be listened on for a panic recovery.
// NOTE that this channel will be closed after the first reported panic.
func (s *PanicSubscriber) PanicSignal() chan struct{} {
	return s.signal
}
