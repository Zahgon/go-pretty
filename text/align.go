package text

// Align denotes how text is to be aligned horizontally.
type Align int

// Align enumerations
const (
	AlignDefault Align = iota // same as AlignLeft
	AlignLeft                 // "left        "
	AlignCenter               // "   center   "
	AlignJustify              // "justify   it"
	AlignRight                // "       right"
	AlignAuto                 // AlignRight for numbers, AlignLeft for the rest
)

// Apply aligns the text as directed. For ex.:
//   - AlignDefault.Apply("Jon Snow", 12) returns "Jon Snow    "
//   - AlignLeft.Apply("Jon Snow",    12) returns "Jon Snow    "
//   - AlignCenter.Apply("Jon Snow",  12) returns "  Jon Snow  "
//   - AlignJustify.Apply("Jon Snow", 12) returns "Jon     Snow"
//   - AlignRight.Apply("Jon Snow",   12) returns "    Jon Snow"
//   - AlignAuto.Apply("Jon Snow",    12) returns "Jon Snow    "
func (a Align) Apply(text string, maxLength int) string { _ = "STUB: not implemented"; return "" }

// was able to parse a number out of the string

// now, align the text

// left pad with half the number of spaces needed before using %text

// HTMLProperty returns the equivalent HTML horizontal-align tag property.
func (a Align) HTMLProperty() string { _ = "STUB: not implemented"; return "" }

// MarkdownProperty returns the equivalent Markdown horizontal-align separator.
// An optional minLength can be provided to extend the dashes to match the
// column content width; the result will be max(minLength, 3)+2 wide (including
// leading/trailing space or colon). Without minLength (or 0), it defaults to 3.
func (a Align) MarkdownProperty(minLength ...int) string { _ = "STUB: not implemented"; return "" }

func (a Align) trimString(text string) string { _ = "STUB: not implemented"; return "" }

func justifyText(text string, textLength int, maxLength int) string {
	_ = "STUB: not implemented"
	// split the text into individual words
	return ""
}

// empty string implies result is just spaces for maxLength

// get the number of spaces to insert into the text

// textLength (display-width) exceeds maxLength; this can happen
// when the cell contains wide Unicode characters (e.g. CJK) whose
// display width is greater than their rune count. Return the text
// as-is; truncation is the caller's responsibility.

// create the output string word by word with spaces in between

// insert spaces only after the first word

// insert all the remaining space before the last word

// insert the determined number of spaces between each word

// and reduce the number of spaces needed after this
