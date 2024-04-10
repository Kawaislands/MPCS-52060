// Lamont Samuels 04/18/19
// atomicbool.go

// Package ppsync provides basic mutual exclusion lock primitives.
package ppsync

import "sync/atomic"

// Representation for a boolean integer value
const (
	TRUE  = 1
	FALSE = 0
)

//AtomicBool represents an atomic boolean type in Go
type AtomicBool struct {
	value int32
}

func cvtBool(value bool) int32 {

	if value {
		return TRUE
	}
	return FALSE
}
func cvtInt(value int32) bool {

	if value == TRUE {
		return true
	}
	return false
}
// NewABool creates and initializes an atomic
// boolean
func NewABool(initValue bool) AtomicBool {

	return AtomicBool{cvtBool(initValue)}
}

// GetAndSet performs an atomic swap operation on
// the passed in the boolean value
func (aBool *AtomicBool) GetAndSet(newValue bool) bool {
	old := atomic.SwapInt32(&aBool.value, cvtBool(newValue))
	return cvtInt(old)
}

// Get returns the current value. This is not atomic
func (aBool *AtomicBool) Get() bool {
	return cvtInt(aBool.value)
}

// Set sets the current value. This is not atomic
func (aBool *AtomicBool) Set(newValue bool) {
	aBool.value = cvtBool(newValue)
}
