package progress

import "sort"

// SortBy helps sort a list of Trackers by various means.
type SortBy int

const (
	// SortByNone doesn't do any sorting == sort by insertion order.
	SortByNone SortBy = iota

	// SortByIndex sorts by the Index field in ascending order. Index 0 comes
	// first. To reorder done trackers as well, set
	// StyleOptions.KeepTrackersTogether = true; otherwise done trackers stay
	// frozen in completion order and only active trackers are sorted by Index.
	SortByIndex

	// SortByIndexDsc sorts by the Index field in descending order. Higher
	// indices come first. To reorder done trackers as well, set
	// StyleOptions.KeepTrackersTogether = true; otherwise done trackers stay
	// frozen in completion order and only active trackers are sorted by Index.
	SortByIndexDsc

	// SortByMessage sorts by the Message alphabetically in ascending order.
	SortByMessage

	// SortByMessageDsc sorts by the Message alphabetically in descending order.
	SortByMessageDsc

	// SortByPercent sorts by the Percentage complete in ascending order.
	SortByPercent

	// SortByPercentDsc sorts by the Percentage complete in descending order.
	SortByPercentDsc

	// SortByValue sorts by the Value in ascending order.
	SortByValue

	// SortByValueDsc sorts by the Value in descending order.
	SortByValueDsc
)

// Sort applies the sorting method defined by SortBy.
func (sb SortBy) Sort(trackers []*Tracker) { _ = "STUB: not implemented"; return }

// no sort

type sortByIndex []*Tracker

func (sb sortByIndex) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByIndex) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByIndex) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Same index: maintain insertion order (use timeStart as tiebreaker)

type sortByIndexDsc []*Tracker

func (sb sortByIndexDsc) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByIndexDsc) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByIndexDsc) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Same index: maintain insertion order (earlier timeStart first)

// Reverse: higher index comes first

type sortByMessage []*Tracker

func (sb sortByMessage) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByMessage) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByMessage) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type sortByPercent []*Tracker

func (sb sortByPercent) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByPercent) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByPercent) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// When percentages are equal, preserve insertion order (earlier timeStart first)

type sortByPercentDsc []*Tracker

func (sb sortByPercentDsc) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByPercentDsc) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByPercentDsc) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// When percentages are equal, preserve insertion order (earlier timeStart first)

// Reverse: higher percentage comes first

type sortByValue []*Tracker

func (sb sortByValue) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByValue) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByValue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type sortByValueDsc []*Tracker

func (sb sortByValueDsc) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sb sortByValueDsc) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sb sortByValueDsc) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// When values are equal, preserve insertion order (earlier timeStart first)

// Reverse: higher value comes first

type sortDsc struct{ sort.Interface }

func (sd sortDsc) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// Reverse the comparison for descending order
	// When elements are equal (both Less calls return false), preserve insertion order
	return false
}
