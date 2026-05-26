//go:build windows
// +build windows

package text

import (
	"sync"
)

var enableVTPMutex = sync.Mutex{}

func areANSICodesSupported() bool { _ = "STUB: not implemented"; return false }
