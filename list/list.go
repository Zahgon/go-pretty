package list

import (
	"io"
	"strings"
)

const (
	// DefaultHTMLCSSClass stores the css-class to use when none-provided via
	// SetHTMLCSSClass(cssClass string).
	DefaultHTMLCSSClass = "go-pretty-table"
)

// listItem represents one line in the List
type listItem struct {
	Level int
	Text  string
}

// List helps print a 2-dimensional array in a human-readable pretty-List.
type List struct {
	// approxSize stores the approximate output length/size
	approxSize int
	// htmlCSSClass stores the HTML CSS Class to use on the <ul> node
	htmlCSSClass string
	// items contains the list of items to render
	items []*listItem
	// level stores the current indentation level
	level int
	// outputMirror stores an io.Writer where the "Render" functions would write
	outputMirror io.Writer
	// style contains all the strings used to draw the List, and more
	style *Style
}

// AppendItem appends the item to the List of items to render.
func (l *List) AppendItem(item interface{}) { _ = "STUB: not implemented"; return }

// AppendItems appends the items to the List of items to render.
func (l *List) AppendItems(items []interface{}) { _ = "STUB: not implemented"; return }

// Indent indents the following items to appear right-shifted.
func (l *List) Indent() { _ = "STUB: not implemented"; return }

// should not indent when there is no item in the current level

// already indented compared to previous item; do not indent more

// Length returns the number of items to be rendered.
func (l *List) Length() int { _ = "STUB: not implemented"; return 0 }

// Reset sets the List to its initial state.
func (l *List) Reset() { _ = "STUB: not implemented"; return }

// SetHTMLCSSClass sets the HTML CSS Class to use on the <ul> node
// when rendering the List in HTML format. Recursive lists would use a numbered
// index suffix. For ex., if the cssClass is set as "foo"; the <ul> for level 0
// would have the class set as "foo"; the <ul> for level 1 would have "foo-1".
func (l *List) SetHTMLCSSClass(cssClass string) { _ = "STUB: not implemented"; return }

// SetOutputMirror sets an io.Writer for all the Render functions to "Write" to
// in addition to returning a string.
func (l *List) SetOutputMirror(mirror io.Writer) { _ = "STUB: not implemented"; return }

// SetStyle overrides the DefaultStyle with the provided one.
func (l *List) SetStyle(style Style) {
	_ = "STUB: not implemented"

	// Style returns the current style.
	return
}

func (l *List) Style() *Style { _ = "STUB: not implemented"; return nil }

func (l *List) analyzeAndStringify(item interface{}) *listItem {
	_ = "STUB: not implemented"
	return nil
}

// UnIndent un-indents the following items to appear left-shifted.
func (l *List) UnIndent() { _ = "STUB: not implemented"; return }

func (l *List) UnIndentAll() { _ = "STUB: not implemented"; return }

func (l *List) initForRender() {
	_ = "STUB: not implemented"
	// pick a default style
	return
}

// calculate the approximate size needed by looking at all entries

// account for the following when incrementing approxSize:
// 1. prefix, 2. padding, 3. bullet, 4. text, 5. newline

// default to a HTML CSS Class if none-defined

func (l *List) hasMoreItemsInLevel(levelIdx int, fromItemIdx int) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *List) render(out *strings.Builder) string { _ = "STUB: not implemented"; return "" }

// renderHint has hints for the Render*() logic
type renderHint struct {
	isTopItem    bool
	isFirstItem  bool
	isOnlyItem   bool
	isLastItem   bool
	isBottomItem bool
}
