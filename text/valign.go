package text

// VAlign denotes how text is to be aligned vertically.
type VAlign int

// VAlign enumerations
const (
	VAlignDefault VAlign = iota // same as VAlignTop
	VAlignTop                   // "top\n\n"
	VAlignMiddle                // "\nmiddle\n"
	VAlignBottom                // "\n\nbottom"
)

// Apply aligns the lines vertically. For ex.:
//   - VAlignTop.Apply({"Game", "Of", "Thrones"},    5)
//     returns {"Game", "Of", "Thrones", "", ""}
//   - VAlignMiddle.Apply({"Game", "Of", "Thrones"}, 5)
//     returns {"", "Game", "Of", "Thrones", ""}
//   - VAlignBottom.Apply({"Game", "Of", "Thrones"}, 5)
//     returns {"", "", "Game", "Of", "Thrones"}
func (va VAlign) Apply(lines []string, maxLines int) []string {
	_ = "STUB: not implemented"
	return nil
}

// ApplyStr aligns the string (of 1 or more lines) vertically. For ex.:
//   - VAlignTop.ApplyStr("Game\nOf\nThrones",    5)
//     returns {"Game", "Of", "Thrones", "", ""}
//   - VAlignMiddle.ApplyStr("Game\nOf\nThrones", 5)
//     returns {"", "Game", "Of", "Thrones", ""}
//   - VAlignBottom.ApplyStr("Game\nOf\nThrones", 5)
//     returns {"", "", "Game", "Of", "Thrones"}
func (va VAlign) ApplyStr(text string, maxLines int) []string {
	_ = "STUB: not implemented"
	return nil
}

// HTMLProperty returns the equivalent HTML vertical-align tag property.
func (va VAlign) HTMLProperty() string { _ = "STUB: not implemented"; return "" }
