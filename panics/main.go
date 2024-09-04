package main

import (
	"fmt"

	"github.com/ellemouton/playgrounds/panics/fn"
	"github.com/ellemouton/playgrounds/panics/loop"
	"github.com/ellemouton/playgrounds/panics/pool"
)

type LiT struct {
	Loop      *loop.Loop
	loopPanic *fn.PanicSubscriber

	Pool      *pool.Pool
	poolPanic *fn.PanicSubscriber
}

func NewLiT() *LiT {
	return &LiT{
		Loop:      loop.New(),
		Pool:      pool.New(),
		loopPanic: fn.NewPanicSubscriber(),
		poolPanic: fn.NewPanicSubscriber(),
	}
}

func (l *LiT) StartWithOutRecovery() {
	l.Loop.StartAsSubserver(nil)
	l.Pool.StartAsSubserver(nil)
}

func (l *LiT) StartWithRecovery() {
	var ()

	l.Loop.StartAsSubserver(l.loopPanic)
	l.Pool.StartAsSubserver(l.poolPanic)
}

func (l *LiT) Stop() {
	l.Loop.Stop()
	l.Pool.Stop()
}

func main() {
	// Example 1: Start loop with no panic subscriber. This will cause the
	// panic ot be re-thrown.
	//example1()

	// Example 2: Start loop with a panic subscriber. Should catch it and
	// gracefully shutdown. (NOTE: comment out example 1!)
	example2()
}

func example1() {
	lit := NewLiT()

	lit.StartWithOutRecovery()

	lit.Stop()
}

func example2() {
	lit := NewLiT()

	lit.StartWithRecovery()

	lit.Pool.TriggerPanicNow()

	lit.Stop()

	select {
	case <-lit.loopPanic.PanicSignal():
		fmt.Println("loop did panic", lit.loopPanic.PanicLogs())
	default:
	}

	select {
	case <-lit.poolPanic.PanicSignal():
		fmt.Println("pool did panic", lit.poolPanic.PanicLogs())
	default:
	}

	fmt.Println("Gracefully shutting down")
}
