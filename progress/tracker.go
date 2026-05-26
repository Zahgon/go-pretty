package progress

import (
	"sync"
	"time"
)

// Tracker helps track the progress of a single task. The way to use it is to
// instantiate a Tracker with a valid Message, a valid (expected) Total, and
// Units values. This should then be fed to the Progress Writer with the
// Writer.AppendTracker() method. When the task that is being done has progress,
// increment the value using the Tracker.Increment(value) method.
type Tracker struct {
	// AutoStopDisabled prevents the tracker from marking itself as done when
	// the value goes beyond the total (if set). Note that this means that a
	// manual call to MarkAsDone or MarkAsErrored is expected.
	AutoStopDisabled bool
	// DeferStart prevents the tracker from starting immediately when appended.
	// It will be rendered but remain dormant until Start, Increment,
	// IncrementWithError or SetValue is called.
	DeferStart bool
	// ExpectedDuration tells how long this task is expected to take; and will
	// be used in calculation of the ETA value
	ExpectedDuration time.Duration
	// Index specifies the explicit order for this tracker. When SortByIndex
	// is used, trackers are sorted by this value regardless of completion status.
	// Lower values appear first, with 0 being the first index.
	Index uint64
	// Message should contain a short description of the "task"; please note
	// that this should NOT be updated in the middle of progress - you should
	// instead use UpdateMessage() to do this safely without hitting any race
	// conditions
	Message string
	// RemoveOnCompletion tells the Progress Bar to remove this tracker when
	// it is done, instead of rendering a "completed" line
	RemoveOnCompletion bool
	// Total should be set to the (expected) Total/Final value to be reached
	Total int64
	// Units defines the type of the "value" being tracked
	Units Units

	done      bool
	err       bool
	mutex     sync.RWMutex
	timeStart time.Time
	timeStop  time.Time
	value     int64
	minETA    time.Duration
}

// ETA returns the expected time of "arrival" or completion of this tracker. It
// is an estimate and is not guaranteed.
func (t *Tracker) ETA() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Increment updates the current value of the task being tracked.
func (t *Tracker) Increment(value int64) { _ = "STUB: not implemented"; return }

// IncrementWithError updates the current value of the task being tracked and
// marks that an error occurred.
func (t *Tracker) IncrementWithError(value int64) { _ = "STUB: not implemented"; return }

// IsStarted true if the tracker has started, false when using DeferStart
// prior to Start, Increment, IncrementWithError or SetValue being called.
func (t *Tracker) IsStarted() bool { _ = "STUB: not implemented"; return false }

// IsDone returns true if the tracker is done (value has reached the expected
// Total set during initialization).
func (t *Tracker) IsDone() bool { _ = "STUB: not implemented"; return false }

// IsErrored true if an error was set with IncrementWithError or MarkAsErrored.
func (t *Tracker) IsErrored() bool { _ = "STUB: not implemented"; return false }

// IsIndeterminate returns true if the tracker is indeterminate; i.e., the total
// is unknown and it is impossible to auto-calculate if tracking is done.
func (t *Tracker) IsIndeterminate() bool { _ = "STUB: not implemented"; return false }

// MarkAsDone forces completion of the tracker by updating the current value as
// the expected Total value.
func (t *Tracker) MarkAsDone() { _ = "STUB: not implemented"; return }

// MarkAsErrored forces completion of the tracker by updating the current value as
// the expected Total value, and recording as error.
func (t *Tracker) MarkAsErrored() {
	_ = "STUB: not implemented"

	// only update error if not done and if not previously set
	return
}

// PercentDone returns the currently completed percentage value.
func (t *Tracker) PercentDone() float64 { _ = "STUB: not implemented"; return 0 }

func (t *Tracker) percentDoneWithoutLock() float64 { _ = "STUB: not implemented"; return 0 }

// Reset resets the tracker to its initial state.
func (t *Tracker) Reset() { _ = "STUB: not implemented"; return }

// SetValue sets the value of the tracker and re-calculates if the tracker is
// "done".
func (t *Tracker) SetValue(value int64) { _ = "STUB: not implemented"; return }

// Start starts the tracking for the case when DeferStart=false.
func (t *Tracker) Start() { _ = "STUB: not implemented"; return }

// UpdateMessage updates the message string.
func (t *Tracker) UpdateMessage(msg string) { _ = "STUB: not implemented"; return }

// UpdateTotal updates the total value.
func (t *Tracker) UpdateTotal(total int64) { _ = "STUB: not implemented"; return }

// Value returns the current value of the tracker.
func (t *Tracker) Value() int64 { _ = "STUB: not implemented"; return 0 }

func (t *Tracker) incrementWithoutLock(value int64) { _ = "STUB: not implemented"; return }

func (t *Tracker) message() string { _ = "STUB: not implemented"; return "" }

func (t *Tracker) start() { _ = "STUB: not implemented"; return }

func (t *Tracker) startWithoutLock() { _ = "STUB: not implemented"; return }

// this must be called with the mutex held with a write lock
func (t *Tracker) stop() { _ = "STUB: not implemented"; return }

func (t *Tracker) valueAndTotal() (int64, int64) { _ = "STUB: not implemented"; return 0, 0 }

// timeStartStopAndDone returns timeStart, timeStop, and done under a single
// RLock so callers see a consistent snapshot — needed because an active→done
// transition between separate reads would yield a non-zero timeStart with a
// zero timeStop while done is true, breaking timeStop.Sub(timeStart).
func (t *Tracker) timeStartStopAndDone() (time.Time, time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), false
}

// timeStartValue returns the start time safely.
func (t *Tracker) timeStartValue() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
