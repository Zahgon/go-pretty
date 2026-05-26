package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	flagAutoStop           = flag.Bool("auto-stop", false, "Auto-stop rendering?")
	flagCustomRenderer     = flag.Bool("custom-renderer", false, "Use custom render functions with rainbow colors")
	flagHideETA            = flag.Bool("hide-eta", false, "Hide the ETA?")
	flagHideETAOverall     = flag.Bool("hide-eta-overall", false, "Hide the ETA in the overall tracker?")
	flagHideOverallTracker = flag.Bool("hide-overall", false, "Hide the Overall Tracker?")
	flagHidePercentage     = flag.Bool("hide-percentage", false, "Hide the progress percent?")
	flagHideTime           = flag.Bool("hide-time", false, "Hide the time taken?")
	flagHideValue          = flag.Bool("hide-value", false, "Hide the tracker value?")
	flagNumTrackers        = flag.Int("num-trackers", 13, "Number of Trackers")
	flagRandomFail         = flag.Bool("rnd-fail", false, "Introduce random failures in tracking")
	flagRandomDefer        = flag.Bool("rnd-defer", false, "Introduce random deferred starts")
	flagRandomRemove       = flag.Bool("rnd-remove", false, "Introduce random remove of trackers on completion")
	flagRandomLogs         = flag.Bool("rnd-logs", false, "Output random logs in the middle of tracking")
	flagSortBy             = flag.String("sort-by", "percent-dsc", "Sort trackers by? (none, index, index-dsc, message, message-dsc, percent, percent-dsc, value, value-dsc)")
	flagShowSpeed          = flag.Bool("show-speed", false, "Show the tracker speed?")
	flagShowSpeedOverall   = flag.Bool("show-speed-overall", false, "Show the overall tracker speed?")
	flagShowPinned         = flag.Bool("show-pinned", false, "Show a pinned message?")

	messageColors = []text.Color{
		text.FgRed,
		text.FgGreen,
		text.FgYellow,
		text.FgBlue,
		text.FgMagenta,
		text.FgCyan,
		text.FgWhite,
	}
	rng       = rand.New(rand.NewSource(time.Now().UnixNano()))
	timeStart = time.Now()
)

// customTrackerDeterminateRenderer creates a progress bar using rainbow colors for determinate progress
func customTrackerDeterminateRenderer(value int64, total int64, maxLen int) string {
	_ = "STUB: not implemented"
	return ""
}

// Use rainbow colors based on position in the progress bar
// Map position to 6 rainbow colors

// customTrackerIndeterminateRenderer creates a progress bar using rotating rainbow colors for indeterminate progress
func customTrackerIndeterminateRenderer(maxLen int) string {
	_ = "STUB: not implemented"
	// For indeterminate progress, use rotating rainbow colors
	return ""
}

func getMessage(idx int64, units *progress.Units) string { _ = "STUB: not implemented"; return "" }

func getSortBy() progress.SortBy { _ = "STUB: not implemented"; return *new(progress.SortBy) }

func getUnits(idx int64) *progress.Units { _ = "STUB: not implemented"; return nil }

func trackSomething(pw progress.Writer, idx int64, updateMessage bool) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()
	fmt.Printf("Tracking Progress of %d trackers ...\n\n", *flagNumTrackers)

	// instantiate a Progress Writer and set up the options
	pw := progress.NewWriter()
	pw.SetAutoStop(*flagAutoStop)
	pw.SetMessageLength(24)
	pw.SetNumTrackersExpected(*flagNumTrackers)
	pw.SetSortBy(getSortBy())
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerLength(25)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"
	pw.Style().Visibility.ETA = !*flagHideETA
	pw.Style().Visibility.ETAOverall = !*flagHideETAOverall
	pw.Style().Visibility.Percentage = !*flagHidePercentage
	pw.Style().Visibility.Speed = *flagShowSpeed
	pw.Style().Visibility.SpeedOverall = *flagShowSpeedOverall
	pw.Style().Visibility.Time = !*flagHideTime
	pw.Style().Visibility.TrackerOverall = !*flagHideOverallTracker
	pw.Style().Visibility.Value = !*flagHideValue
	pw.Style().Visibility.Pinned = *flagShowPinned

	// set up custom render functions if flag is enabled
	if *flagCustomRenderer {
		pw.Style().Renderer.TrackerDeterminate = customTrackerDeterminateRenderer
		pw.Style().Renderer.TrackerIndeterminate = customTrackerIndeterminateRenderer
	}

	// call Render() in async mode; yes we don't have any trackers at the moment
	go pw.Render()

	// add a bunch of trackers with random parameters to demo most of the
	// features available; do this in async too like a client might do (for ex.
	// when downloading a bunch of files in parallel)
	for idx := int64(1); idx <= int64(*flagNumTrackers); idx++ {
		go trackSomething(pw, idx, idx == int64(*flagNumTrackers))

		// in auto-stop mode, the Render logic terminates the moment it detects
		// zero active trackers; but in a manual-stop mode, it keeps waiting and
		// is a good chance to demo trackers being added dynamically while other
		// trackers are active or done
		if !*flagAutoStop {
			time.Sleep(time.Millisecond * 100)
		}
	}

	// wait for one or more trackers to become active (just blind-wait for a
	// second) and then keep watching until Rendering is in progress
	time.Sleep(time.Second)
	messagesLogged := make(map[string]bool)
	for pw.IsRenderInProgress() {
		if *flagRandomLogs && pw.LengthDone()%3 == 0 {
			logMsg := text.Faint.Sprintf("[INFO] done with %d trackers", pw.LengthDone())
			if !messagesLogged[logMsg] {
				pw.Log(logMsg)
				messagesLogged[logMsg] = true
			}
		}

		// for manual-stop mode, stop when there are no more active trackers
		if !*flagAutoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("\nAll done!")
}
