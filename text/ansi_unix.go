//go:build !windows
// +build !windows

package text

func areANSICodesSupported() bool {
	_ = "STUB: not implemented"
	// On Unix systems, ANSI codes are generally supported unless TERM is "dumb"
	// This is a basic check; 256-color sequences are ANSI sequences and will
	// be handled by terminals that support them (or ignored by those that don't)
	return false
}
