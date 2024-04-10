// Lamont Samuels 04/18/19
// taslock.go

// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import (
	"runtime"
	"sync"
)

// A TASLock is a mutual exclusion lock that represents a test-and-set lock.
// The zero value for a TASLock is an unlocked mutex.
type TASLock struct {
	state *AtomicBool
}

// Creates and returns a new TASLock.
func NewTASLock() sync.Locker {
	atomicBool := NewABool(false)
	return &TASLock{&atomicBool}
}

// Lock locks lock. If the lock is already in use, the calling goroutine
// blocks until the lock is available.
func (lock *TASLock) Lock() {
	for lock.state.GetAndSet(true) {
		runtime.Gosched()
	}
}

// Unlock unlocks lock.
// It is a run-time error if lock is not locked on entry to Unlock.
//
// A locked lock is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a lock and then
// arrange for another goroutine to unlock it.
func (lock *TASLock) Unlock() {
	lock.state.Set(false)
}
