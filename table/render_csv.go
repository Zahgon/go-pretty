package table

import (
	"strings"
)

// RenderCSV renders the Table in CSV format. Example:
//
//	#,First Name,Last Name,Salary,
//	1,Arya,Stark,3000,
//	20,Jon,Snow,2000,"You know nothing\, Jon Snow!"
//	300,Tyrion,Lannister,5000,
//	,,Total,10000,
func (t *Table) RenderCSV() string { _ = "STUB: not implemented"; return "" }

func (t *Table) csvFixCommas(str string) string { _ = "STUB: not implemented"; return "" }

func (t *Table) csvFixDoubleQuotes(str string) string { _ = "STUB: not implemented"; return "" }

func (t *Table) csvRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	// when working on line number 2 or more, insert a newline first
	return
}

// generate the columns to render in CSV format and append to "out"

// auto-index column

func (t *Table) csvRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	return
}
