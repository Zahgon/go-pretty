package table

import (
	"strings"
)

// RenderMarkdown renders the Table in Markdown format. Example:
//
//	| # | First Name | Last Name | Salary |  |
//	| ---:| --- | --- | ---:| --- |
//	| 1 | Arya | Stark | 3000 |  |
//	| 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
//	| 300 | Tyrion | Lannister | 5000 |  |
//	|  |  | Total | 10000 |  |
func (t *Table) RenderMarkdown() string { _ = "STUB: not implemented"; return "" }

func (t *Table) markdownRenderCaption(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) markdownRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	// when working on line number 2 or more, insert a newline first
	return
}

// render each column up to the max. columns seen in all the rows

func (t *Table) markdownRenderRowAutoIndex(out *strings.Builder, colIdx int, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) markdownRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) markdownRenderRowsFooter(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) markdownRenderRowsHeader(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) markdownRenderSeparator(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	// when working on line number 2 or more, insert a newline first
	return
}

func (t *Table) markdownRenderTitle(out *strings.Builder) { _ = "STUB: not implemented"; return }
