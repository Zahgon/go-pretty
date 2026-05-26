package table

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	// DefaultHTMLCSSClass stores the css-class to use when none-provided via
	// SetHTMLCSSClass(cssClass string).
	DefaultHTMLCSSClass = "go-pretty-table"
)

// RenderHTML renders the Table in HTML format. Example:
//
//	<table class="go-pretty-table">
//	  <thead>
//	  <tr>
//	    <th align="right">#</th>
//	    <th>First Name</th>
//	    <th>Last Name</th>
//	    <th align="right">Salary</th>
//	    <th>&nbsp;</th>
//	  </tr>
//	  </thead>
//	  <tbody>
//	  <tr>
//	    <td align="right">1</td>
//	    <td>Arya</td>
//	    <td>Stark</td>
//	    <td align="right">3000</td>
//	    <td>&nbsp;</td>
//	  </tr>
//	  <tr>
//	    <td align="right">20</td>
//	    <td>Jon</td>
//	    <td>Snow</td>
//	    <td align="right">2000</td>
//	    <td>You know nothing, Jon Snow!</td>
//	  </tr>
//	  <tr>
//	    <td align="right">300</td>
//	    <td>Tyrion</td>
//	    <td>Lannister</td>
//	    <td align="right">5000</td>
//	    <td>&nbsp;</td>
//	  </tr>
//	  </tbody>
//	  <tfoot>
//	  <tr>
//	    <td align="right">&nbsp;</td>
//	    <td>&nbsp;</td>
//	    <td>Total</td>
//	    <td align="right">10000</td>
//	    <td>&nbsp;</td>
//	  </tr>
//	  </tfoot>
//	</table>
func (t *Table) RenderHTML() string { _ = "STUB: not implemented"; return "" }

func (t *Table) htmlGetColStrAndTag(row rowStr, colIdx int, hint renderHint) (string, string) {
	_ = "STUB: not implemented"
	// get the column contents
	return "", ""
}

// header uses "th" instead of "td"

func (t *Table) htmlRenderCaption(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) htmlRenderColumn(out *strings.Builder, colStr string) {
	_ = "STUB: not implemented"
	// convertEscSequencesToSpans already escapes text content, so skip
	// EscapeText if ConvertColorsToSpans is true
	return
}

func (t *Table) htmlRenderColumnAttributes(out *strings.Builder, colIdx int, hint renderHint, alignOverride text.Align) {
	_ = "STUB: not implemented"
	// determine the HTML "align"/"valign" property values
	return
}

// determine the HTML "class" property values for the colors

func (t *Table) htmlRenderColumnAutoIndex(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) htmlRenderRow(out *strings.Builder, row rowStr, hint renderHint) {
	_ = "STUB: not implemented"
	return
}

// auto-index column

// auto-merged columns should be skipped

// get the real row to consider all lines in each column instead of just
// looking at the current "line"

// write the row

func (t *Table) htmlRenderRows(out *strings.Builder, rows []rowStr, hint renderHint) {
	_ = "STUB: not implemented"

	// determine that tag to use based on the type of the row
	return
}

func (t *Table) htmlRenderRowsFooter(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) htmlRenderRowsHeader(out *strings.Builder) { _ = "STUB: not implemented"; return }

func (t *Table) htmlRenderTitle(out *strings.Builder) { _ = "STUB: not implemented"; return }
