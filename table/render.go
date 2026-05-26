package table

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Render renders the Table in a human-readable "pretty" format. Example:
//
//	┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
//	│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
//	├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//	│   1 │ Arya       │ Stark     │   3000 │                             │
//	│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
//	│ 300 │ Tyrion     │ Lannister │   5000 │                             │
//	├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
//	│     │            │ TOTAL     │  10000 │                             │
//	└─────┴────────────┴───────────┴────────┴─────────────────────────────┘
func (t *Table) Render() string { _ = "STUB: not implemented"; return "" }

// top-most border

// header rows

// (data) rows

// footer rows

// bottom-most border

// caption

//gocyclo:ignore
func (t *Table) renderColumn(out *strings.Builder, row rowStr, colIdx int, maxColumnLength int, hint renderHint) int {
	_ = "STUB: not implemented"
	return 0

	// when working on the first column, and autoIndex is true, insert a new
	// column with the row number on it.
}

// when working on column number 2 or more, render the column separator

// extract the text, convert-case if not-empty and align horizontally

// leave colStr empty; align will expand the column as necessary

// if horizontal cell merges are enabled, look ahead and see how many cells
// have the same content and merge them all until a cell with a different
// content is found; override alignment to Center in this case

// get the real row to consider all lines in each column instead of just
// looking at the current "line"

// pad both sides of the column

func (t *Table) renderColumnAutoIndex(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderColumnColorized(out *strings.Builder, colIdx int, colStr string, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderColumnSeparator(out *strings.Builder, row rowStr, colIdx int, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderLine(out *strings.Builder, row rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	// if the output has content, it means that this call is working on line
	// number 2 or more; separate them with a newline
	return
}

// use a brand-new strings.Builder if a row length limit has been set

// grow the strings.Builder to the maximum possible row length

// merge the strings.Builder objects if a new one was created earlier

// if a page size has been set, and said number of lines has already
// been rendered, and the header is not being rendered right now, render
// the header all over again with a spacing line

func (t *Table) renderLineMergeOutputs(out *strings.Builder, outLine *strings.Builder) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderMarginLeft(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderMarginRight(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderRow(out *strings.Builder, row rowStr, hint renderHint) {
	_ = "STUB: not implemented"

	// fit every column into the allowedColumnLength/maxColumnLength limit
	// and in the process find the max. number of lines in any column in
	// this row
	return
}

// if there is just 1 line in all columns, add the row as such; else
// split each column into individual lines and render them one-by-one

// convert one row into N # of rows based on colMaxLines

func (t *Table) renderRowSeparator(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) renderRowsBorderBottom(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) renderRowsBorderTop(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) renderRowsFooter(out *strings.Builder) { _ = "STUB: not implemented"; return }

// Only add separator before footer if there are data rows.
// Otherwise, renderRowsHeader already added one.

func (t *Table) renderRowsHeader(out *strings.Builder) { _ = "STUB: not implemented"; return }

// Only add separator after header if there are data rows or footer rows.
// Otherwise, the bottom border is rendered directly.

func (t *Table) renderTitle(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) renderTitleLine(out *strings.Builder, lenText int, titleLine string, colors text.Colors, colorsBorder text.Colors) {
	_ = "STUB: not implemented"
	return
}
