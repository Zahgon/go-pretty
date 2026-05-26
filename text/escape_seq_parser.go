package text

// Constants
const (
	EscapeReset        = EscapeResetCSI
	EscapeResetCSI     = EscapeStartCSI + "0" + EscapeStopCSI
	EscapeResetOSI     = EscapeStartOSI + "0" + EscapeStopOSI
	EscapeStart        = EscapeStartCSI
	EscapeStartCSI     = "\x1b["
	EscapeStartOSI     = "\x1b]"
	EscapeStartRune    = rune(27) // \x1b
	EscapeStartRuneCSI = '['      // [
	EscapeStartRuneOSI = ']'      // ]
	EscapeStop         = EscapeStopCSI
	EscapeStopCSI      = "m"
	EscapeStopOSI      = "\\"
	EscapeStopRune     = EscapeStopRuneCSI
	EscapeStopRuneCSI  = 'm'
	EscapeStopRuneOSI  = '\\'
)

// Deprecated Constants
const (
	CSIStartRune = EscapeStartRuneCSI
	CSIStopRune  = EscapeStopRuneCSI
	OSIStartRune = EscapeStartRuneOSI
	OSIStopRune  = EscapeStopRuneOSI
)

type escSeqKind int

const (
	escSeqKindUnknown escSeqKind = iota
	escSeqKindCSI
	escSeqKindOSI
)

// private constants
const (
	escCodeResetAll        = 0
	escCodeResetIntensity  = 22
	escCodeResetItalic     = 23
	escCodeResetUnderline  = 24
	escCodeResetBlink      = 25
	escCodeResetReverse    = 27
	escCodeResetCrossedOut = 29
	escCodeBold            = 1
	escCodeDim             = 2
	escCodeItalic          = 3
	escCodeUnderline       = 4
	escCodeBlinkSlow       = 5
	escCodeBlinkRapid      = 6
	escCodeReverse         = 7
	escCodeConceal         = 8
	escCodeCrossedOut      = 9

	// conceal OSI sequences
	escapeStartConcealOSI = "\x1b]8;"
	escapeStopConcealOSI  = "\x1b\\"
)

// 256-color codes
const (
	escCode256FgStart = 38
	escCode256BgStart = 48
	escCode256Color   = 5
	escCodeResetFg    = 39
	escCodeResetBg    = 49
	escCode256Max     = 255
)

// Internal encoding for 256-color codes uses fg256Start and bg256Start from color.go
// Private constants initialized from private constants to avoid repeated casting in hot paths
// Foreground 256-color: fg256Start + colorIndex (1000-1255)
// Background 256-color: bg256Start + colorIndex (2000-2255)
const (
	escCode256FgBase = int(fg256Start) // 1000
	escCode256BgBase = int(bg256Start) // 2000
)

// Standard color code ranges
const (
	// Standard foreground colors (30-37)
	escCodeFgStdStart = 30
	escCodeFgStdEnd   = 37
	// Bright foreground colors (90-97)
	escCodeFgBrightStart = 90
	escCodeFgBrightEnd   = 97
	// Standard background colors (40-47)
	escCodeBgStdStart = 40
	escCodeBgStdEnd   = 47
	// Bright background colors (100-107)
	escCodeBgBrightStart = 100
	escCodeBgBrightEnd   = 107
)

// Special characters
const (
	escRuneBEL = '\a' // BEL character (ASCII 7)
)

// EscSeqParser parses ANSI escape sequences from text and tracks active formatting codes.
// It supports both CSI (Control Sequence Introducer) and OSI (Operating System Command)
// escape sequence formats.
type EscSeqParser struct {
	// codes tracks active escape sequence codes (e.g., 1 for bold, 3 for italic).
	codes map[int]bool

	// inEscSeq indicates whether the parser is currently inside an escape sequence.
	inEscSeq bool
	// escSeqKind identifies the type of escape sequence being parsed (CSI or OSI).
	escSeqKind escSeqKind
	// escapeSeq accumulates the current escape sequence being parsed.
	escapeSeq string
}

func (s *EscSeqParser) Codes() []int { _ = "STUB: not implemented"; return nil }

func (s *EscSeqParser) Consume(char rune) { _ = "STUB: not implemented"; return }

// --- FIX for OSC 8 hyperlinks (e.g. \x1b]8;;url\x07label\x1b]8;;\x07)

// BEL

func (s *EscSeqParser) InSequence() bool { _ = "STUB: not implemented"; return false }

func (s *EscSeqParser) IsOpen() bool { _ = "STUB: not implemented"; return false }

func (s *EscSeqParser) ParseSeq(seq string, seqKind escSeqKind) { _ = "STUB: not implemented"; return }

func (s *EscSeqParser) ParseString(str string) string { _ = "STUB: not implemented"; return "" }

func (s *EscSeqParser) Reset() { _ = "STUB: not implemented"; return }

func (s *EscSeqParser) Sequence() string { _ = "STUB: not implemented"; return "" }

// Check if this is a 256-color foreground code (1000-1255)

// 256-color background code (2000-2255)

// Regular code

// clearAllBackgroundColors clears all background color codes.
func (s *EscSeqParser) clearAllBackgroundColors() { _ = "STUB: not implemented"; return }

// clearAllForegroundColors clears all foreground color codes.
func (s *EscSeqParser) clearAllForegroundColors() { _ = "STUB: not implemented"; return }

// clearColorRange clears standard foreground or background colors.
func (s *EscSeqParser) clearColorRange(isForeground bool) {
	_ = "STUB: not implemented"

	// Clear standard foreground colors (30-37, 90-97)
	return
}

// Clear standard background colors (40-47, 100-107)

func (s *EscSeqParser) isEscapeStopRune(char rune) bool { _ = "STUB: not implemented"; return false }

// isRegularCode checks if a code is a regular code (not a 256-color encoded value).
func (s *EscSeqParser) isRegularCode(codeNum int) bool { _ = "STUB: not implemented"; return false }

// parse256ColorSequence attempts to parse a 256-color sequence starting at index i.
// Returns (colorIndex, base, true) if valid, or (0, 0, false) if not.
func (s *EscSeqParser) parse256ColorSequence(codes []string, i int) (colorIndex int, base int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// process256ColorSequences processes 256-color sequences (38;5;n or 48;5;n) and returns
// a map of indices that were part of valid 256-color sequences.
func (s *EscSeqParser) process256ColorSequences(codes []string) map[int]bool {
	_ = "STUB: not implemented"
	return nil
}

// Skip i+1 and i+2 (loop will increment to i+3)

// processCode handles a single escape code.
func (s *EscSeqParser) processCode(codeNum int) { _ = "STUB: not implemented"; return }

// processRegularCodes processes regular escape codes and reset codes.
func (s *EscSeqParser) processRegularCodes(codes []string, processedIndices map[int]bool) {
	_ = "STUB: not implemented"
	return
}

// set256Color sets a 256-color code and clears conflicting colors.
func (s *EscSeqParser) set256Color(base int, colorIndex int) { _ = "STUB: not implemented"; return }

// Clear other colors in the same range

// splitAndTrimCodes splits the sequence by semicolons and trims whitespace.
func (s *EscSeqParser) splitAndTrimCodes(seq string) []string {
	_ = "STUB: not implemented"
	return nil
}

// stripEscapeSequence removes escape sequence markers from the input string.
func (s *EscSeqParser) stripEscapeSequence(seq string, seqKind escSeqKind) string {
	_ = "STUB: not implemented"
	return ""
}
