package text

// Format lets you transform the text in supported methods while keeping escape
// sequences in the string intact and untouched.
type Format int

// Format enumerations
const (
	FormatDefault Format = iota // default_Case
	FormatLower                 // lower
	FormatTitle                 // Title
	FormatUpper                 // UPPER
)

// Apply converts the text as directed.
func (tc Format) Apply(text string) string { _ = "STUB: not implemented"; return "" }

func toTitle(text string) string { _ = "STUB: not implemented"; return "" }

func toUpper(text string) string { _ = "STUB: not implemented"; return "" }

// isSeparator returns true if the given rune is a separator. This function is
// lifted straight out of the standard library @ strings/strings.go.
func isSeparator(r rune) bool {
	_ = "STUB: not implemented"
	// ASCII alphanumerics and underscore are not separators
	return false
}

// Letters and digits are not separators

// Otherwise, all we can do for now is treat spaces as separators.
