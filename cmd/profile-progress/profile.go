package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/pkg/profile"
)

var (
	tracker1  = progress.Tracker{Message: "Calculating Total   # 1", Total: 1000, Units: progress.UnitsDefault}
	tracker2  = progress.Tracker{Message: "Downloading File    # 2", Total: 1000, Units: progress.UnitsBytes}
	tracker3  = progress.Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: progress.UnitsCurrencyDollar}
	profilers = []func(*profile.Profile){
		profile.CPUProfile,
		profile.MemProfileRate(512),
	}
)

func profileRender(profiler func(profile2 *profile.Profile), n int) {
	_ = "STUB: not implemented"
	return
}

func main() {
	numRenders := 5
	if len(os.Args) > 1 {
		var err error
		numRenders, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid Argument: '%s'\n", os.Args[1])
			os.Exit(1)
		}
	}

	for _, profiler := range profilers {
		profileRender(profiler, numRenders)
	}
}
