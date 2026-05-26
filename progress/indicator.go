package progress

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
)

// IndeterminateIndicator defines the structure for the indicator to indicate
// indeterminate progress. Ex.: {0, <=>}
type IndeterminateIndicator struct {
	Position int
	Text     string
}

// IndeterminateIndicatorGenerator is a function that takes the maximum length
// of the progress bar and returns an IndeterminateIndicator telling the
// indicator string, and the location of the same in the progress bar.
//
// Technically, this could generate and return the entire progress bar string to
// override the full display of the same - this is done by the Dominoes and
// Pac-Man examples below.
type IndeterminateIndicatorGenerator func(maxLen int) IndeterminateIndicator

// IndeterminateIndicatorDominoes simulates a bunch of dominoes falling back and
// forth.
func IndeterminateIndicatorDominoes(duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorMovingBackAndForth incrementally moves from the left to
// right and back for each specified duration. If duration is 0, then every
// single invocation moves the indicator.
func IndeterminateIndicatorMovingBackAndForth(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorMovingLeftToRight incrementally moves from the left to
// right and starts from left again for each specified duration. If duration is
// 0, then every single invocation moves the indicator.
func IndeterminateIndicatorMovingLeftToRight(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorMovingRightToLeft incrementally moves from the right to
// left and starts from right again for each specified duration. If duration is
// 0, then every single invocation moves the indicator.
func IndeterminateIndicatorMovingRightToLeft(indicator string, duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorPacMan simulates a Pac-Man character chomping through
// the progress bar back and forth.
func IndeterminateIndicatorPacMan(duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorColoredDominoes simulates a bunch of colored dominoes falling back and
// forth.
func IndeterminateIndicatorColoredDominoes(duration time.Duration, slashColor, backslashColor text.Color) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// IndeterminateIndicatorPacManChomp simulates a Pac-Man character chomping through the progress
// bar back and forth.
func IndeterminateIndicatorPacManChomp(duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

func indeterminateIndicatorPacManChomp() IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// Alternate between open and closed mouth

func indeterminateIndicatorDominoes() IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	// positive == left to right; negative == right to left
	return *new(IndeterminateIndicatorGenerator)
}

func indeterminateIndicatorMovingBackAndForth(indicator string) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	// positive == left to right; negative == right to left
	return *new(IndeterminateIndicatorGenerator)
}

func indeterminateIndicatorMovingLeftToRight(indicator string) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

func indeterminateIndicatorMovingRightToLeft(indicator string) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

func indeterminateIndicatorPacMan() IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}

// positive == left to right; negative == right to left

// timedIndeterminateIndicatorGenerator ticks based on the given duration. If
// duration is 0, it ticks for every invocation.
func timedIndeterminateIndicatorGenerator(indicatorGenerator IndeterminateIndicatorGenerator, duration time.Duration) IndeterminateIndicatorGenerator {
	_ = "STUB: not implemented"
	return *new(IndeterminateIndicatorGenerator)
}
