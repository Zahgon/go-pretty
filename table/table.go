package table

import (
	"io"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
)

// Table helps print a 2-dimensional array in a human-readable pretty-table.
type Table struct {
	// allowedRowLength is the max allowed length for a row (or line of output)
	allowedRowLength int
	// enable automatic indexing of the rows and columns like a spreadsheet?
	autoIndex bool
	// autoIndexVIndexMaxLength denotes the length in chars for the last row
	autoIndexVIndexMaxLength int
	// caption stores the text to be rendered just below the table; and doesn't
	// get used when rendered as a CSV
	caption string
	// columnIsNonNumeric stores if a column contains non-numbers in all rows
	columnIsNonNumeric []bool
	// columnConfigs stores the custom-configuration for 1 or more columns
	columnConfigs []ColumnConfig
	// columnConfigMap stores the custom-configuration by column
	// number and is generated before rendering
	columnConfigMap map[int]ColumnConfig
	// directionModifier caches the direction modifier string to avoid repeated calls
	directionModifier string
	// firstRowOfPage tells if the renderer is on the first row of a page?
	firstRowOfPage bool
	// htmlCSSClass stores the HTML CSS Class to use on the <table> node
	htmlCSSClass string
	// indexColumn stores the number of the column considered as the "index"
	indexColumn int
	// maxColumnLengths stores the length of the longest line in each column
	maxColumnLengths []int
	// maxMergedColumnLengths stores the longest lengths for merged columns
	// endIndex -> startIndex -> maxMergedLength
	maxMergedColumnLengths map[int]map[int]int
	// maxRowLength stores the length of the longest row
	maxRowLength int
	// numColumns stores the (max.) number of columns seen
	numColumns int
	// numLinesRendered keeps track of the number of lines rendered and helps in
	// paginating long tables
	numLinesRendered int
	// outputMirror stores an io.Writer where the "Render" functions would write
	outputMirror io.Writer
	// pager controls how the output is separated into pages
	pager pager
	// renderMode contains the type of table to render
	renderMode renderMode
	// rows stores the rows that make up the body (in string form)
	rows []rowStr
	// rowsColors stores the text.Colors over-rides for each row as defined by
	// rowPainter or rowPainterWithAttributes
	rowsColors []text.Colors
	// rowsConfigs stores RowConfig for each row
	rowsConfigMap map[int]RowConfig
	// rowsRaw stores the rows that make up the body
	rowsRaw []Row
	// rowsRawFiltered is the filtered version of rowsRaw
	rowsRawFiltered []Row
	// rowsFooter stores the rows that make up the footer (in string form)
	rowsFooter []rowStr
	// rowsFooterConfigs stores RowConfig for each footer row
	rowsFooterConfigMap map[int]RowConfig
	// rowsFooterRaw stores the rows that make up the footer
	rowsFooterRaw []Row
	// rowsHeader stores the rows that make up the header (in string form)
	rowsHeader []rowStr
	// rowsHeaderConfigs stores RowConfig for each header row
	rowsHeaderConfigMap map[int]RowConfig
	// rowsHeaderRaw stores the rows that make up the header
	rowsHeaderRaw []Row
	// rowPainter is a custom function that given a Row, returns the colors to
	// use on the entire row
	rowPainter RowPainter
	// rowPainterWithAttributes is same as rowPainter, but with attributes
	rowPainterWithAttributes RowPainterWithAttributes
	// rowSeparators contains the separator columns (dashes that make up the
	// separators between title/header/body/footer
	rowSeparators map[string]rowStr
	// rowSeparatorStrings contains the separator strings for each separator type
	rowSeparatorStrings map[separatorType]string
	// separators is used to keep track of all rowIndices after which a
	// separator has to be rendered
	separators map[int]bool
	// sortBy stores a map of Column
	sortBy []SortBy
	// sortedRowIndices is the output of sorting
	sortedRowIndices []int
	// filterBy stores the filter criteria
	filterBy []FilterBy
	// style contains all the strings used to draw the table, and more
	style *Style
	// suppressEmptyColumns hides columns which have no content on all regular
	// rows
	suppressEmptyColumns bool
	// suppressTrailingSpaces removes all trailing spaces from the end of the last column
	suppressTrailingSpaces bool
	// title contains the text to appear above the table
	title string
}

// AppendFooter appends the row to the List of footers to render.
//
// Only the first item in the "config" will be tagged against this row.
func (t *Table) AppendFooter(row Row, config ...RowConfig) { _ = "STUB: not implemented"; return }

// AppendHeader appends the row to the List of headers to render.
//
// Only the first item in the "config" will be tagged against this row.
func (t *Table) AppendHeader(row Row, config ...RowConfig) { _ = "STUB: not implemented"; return }

// AppendRow appends the row to the List of rows to render.
//
// Only the first item in the "config" will be tagged against this row.
func (t *Table) AppendRow(row Row, config ...RowConfig) { _ = "STUB: not implemented"; return }

// Keep original rows in sync for filtering

// AppendRows appends the rows to the List of rows to render.
//
// Only the first item in the "config" will be tagged against all the rows.
func (t *Table) AppendRows(rows []Row, config ...RowConfig) { _ = "STUB: not implemented"; return }

// AppendSeparator helps render a separator row after the current last row. You
// could call this function over and over, but it will be a no-op unless you
// call AppendRow or AppendRows in between. Likewise, if the last thing you
// append is a separator, it will not be rendered in addition to the usual table
// separator.
//
// ******************************************************************************
// Please note the following caveats:
//  1. SetPageSize(): this may end up creating consecutive separator rows near
//     the end of a page or at the beginning of a page
//  2. SortBy(): since SortBy could inherently alter the ordering of rows, the
//     separators may not appear after the row it was originally intended to
//     follow
//
// ******************************************************************************
func (t *Table) AppendSeparator() { _ = "STUB: not implemented"; return }

// FilterBy sets the rules for filtering the Rows. All filters are applied with
// AND logic (all must match). Filters are applied before sorting.
func (t *Table) FilterBy(filterBy []FilterBy) { _ = "STUB: not implemented"; return }

// ImportGrid helps import 1d or 2d arrays as rows.
func (t *Table) ImportGrid(grid interface{}) bool { _ = "STUB: not implemented"; return false }

// Length returns the number of rows to be rendered.
func (t *Table) Length() int { _ = "STUB: not implemented"; return 0 }

// Pager returns an object that splits the table output into pages and
// lets you move back and forth through them.
func (t *Table) Pager(opts ...PagerOption) Pager { _ = "STUB: not implemented"; return *new(Pager) }

// use a temporary page separator for splitting up the pages

// backup

// restore on exit

// override

// render

// ResetFooters resets and clears all the Footer rows appended earlier.
func (t *Table) ResetFooters() { _ = "STUB: not implemented"; return }

// ResetHeaders resets and clears all the Header rows appended earlier.
func (t *Table) ResetHeaders() { _ = "STUB: not implemented"; return }

// ResetRows resets and clears all the rows appended earlier.
func (t *Table) ResetRows() { _ = "STUB: not implemented"; return }

// SetAllowedRowLength sets the maximum allowed length or a row (or line of
// output) when rendered as a table. Rows that are longer than this limit will
// be "snipped" to the length. Length has to be a positive value to take effect.
//
// Deprecated: in favor if Style().Size.WidthMax
func (t *Table) SetAllowedRowLength(length int) { _ = "STUB: not implemented"; return }

// SetAutoIndex adds a generated header with columns such as "A", "B", "C", etc.
// and a leading column with the row number similar to what you'd see on any
// spreadsheet application. NOTE: Appending a Header will void this
// functionality.
func (t *Table) SetAutoIndex(autoIndex bool) { _ = "STUB: not implemented"; return }

// SetCaption sets the text to be rendered just below the table. This will not
// show up when the Table is rendered as a CSV.
func (t *Table) SetCaption(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// SetColumnConfigs sets the configs for each Column.
func (t *Table) SetColumnConfigs(configs []ColumnConfig) { _ = "STUB: not implemented"; return }

// SetHTMLCSSClass sets the HTML CSS Class to use on the <table> node
// when rendering the Table in HTML format.
//
// Deprecated: in favor of Style().HTML.CSSClass
func (t *Table) SetHTMLCSSClass(cssClass string) { _ = "STUB: not implemented"; return }

// SetIndexColumn sets the given Column # as the column that has the row
// "Number". Valid values range from 1 to N. Note that this is not 0-indexed.
func (t *Table) SetIndexColumn(colNum int) { _ = "STUB: not implemented"; return }

// SetOutputMirror sets an io.Writer for all the Render functions to "Write" to
// in addition to returning a string.
func (t *Table) SetOutputMirror(mirror io.Writer) { _ = "STUB: not implemented"; return }

// SetPageSize sets the maximum number of lines to render before rendering the
// header rows again. This can be useful when dealing with tables containing a
// long list of rows that can span pages. Please note that the pagination logic
// will not consider Header/Footer lines for paging.
func (t *Table) SetPageSize(numLines int) { _ = "STUB: not implemented"; return }

// SetRowPainter sets up the function which determines the colors to use on a
// row. Before rendering, this function is invoked on all rows and the color
// of each row is determined. This color takes precedence over other ways to
// set color (ColumnConfig.Color*, SetColor*()).
func (t *Table) SetRowPainter(painter interface{}) {
	_ = "STUB: not implemented"
	// TODO: fix interface on major version bump to accept only
	// one type of RowPainter: RowPainterWithAttributes renamed to RowPainter
	return
}

// reset both so only one is set at any given time

// if called as SetRowPainter(RowPainter(func...))

// if called as SetRowPainter(func...)

// SetStyle overrides the DefaultStyle with the provided one.
func (t *Table) SetStyle(style Style) {
	_ = "STUB: not implemented"

	// SetTitle sets the title text to be rendered above the table.
	return
}

func (t *Table) SetTitle(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// SortBy sets the rules for sorting the Rows in the order specified. i.e., the
// first SortBy instruction takes precedence over the second and so on. Any
// duplicate instructions on the same column will be discarded while sorting.
func (t *Table) SortBy(sortBy []SortBy) {
	_ = "STUB: not implemented"

	// Style returns the current style.
	return
}

func (t *Table) Style() *Style { _ = "STUB: not implemented"; return nil }

// override WidthMax with allowedRowLength until allowedRowLength is
// removed from code

// SuppressEmptyColumns hides columns when the column is empty in ALL the
// regular rows.
func (t *Table) SuppressEmptyColumns() { _ = "STUB: not implemented"; return }

// SuppressTrailingSpaces removes all trailing spaces from the output.
func (t *Table) SuppressTrailingSpaces() { _ = "STUB: not implemented"; return }

// calculateNumColumnsFromRaw calculates the number of columns from raw rows and headers
func (t *Table) calculateNumColumnsFromRaw() {
	_ = "STUB: not implemented"

	// Check headers first
	return
}

// Check data rows

// Check footer rows

func (t *Table) getAlign(colIdx int, hint renderHint) text.Align {
	_ = "STUB: not implemented"
	return *new(text.Align)
}

func (t *Table) getAutoIndexColumnIDs() rowStr { _ = "STUB: not implemented"; return *new(rowStr) }

func (t *Table) getBorderColors(hint renderHint) text.Colors {
	_ = "STUB: not implemented"
	return *new(text.Colors)
}

func (t *Table) getBorderLeft(hint renderHint) string { _ = "STUB: not implemented"; return "" }

func (t *Table) getBorderRight(hint renderHint) string { _ = "STUB: not implemented"; return "" }

func (t *Table) getColumnColors(colIdx int, hint renderHint) text.Colors {
	_ = "STUB: not implemented"
	return *new(text.Colors)
}

func (t *Table) getColumnColorsForBorderOrSeparator(hint renderHint) text.Colors {
	_ = "STUB: not implemented"
	return *new(text.Colors)
}

// not nil to force caller to paint with no colors

func (t *Table) getColumnSeparator(row rowStr, colIdx int, hint renderHint) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Table) getColumnSeparatorNonBorder(mergeCellsAbove bool, mergeCellsBelow bool, colIdx int, hint renderHint) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Table) getColumnSeparatorNonBorderAutoIndex(mergeNextCol bool, hint renderHint) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Table) getColumnSeparatorNonBorderNonAutoIndex(mergeCellsAbove bool, mergeCellsBelow bool, mergeCurrCol bool, mergeNextCol bool, hint renderHint) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Table) getColumnTransformer(colIdx int, hint renderHint) text.Transformer {
	_ = "STUB: not implemented"
	return *new(text.Transformer)
}

func (t *Table) getColumnWidthMax(colIdx int) int { _ = "STUB: not implemented"; return 0 }

func (t *Table) getColumnWidthMin(colIdx int) int { _ = "STUB: not implemented"; return 0 }

func (t *Table) getFormat(hint renderHint) text.Format {
	_ = "STUB: not implemented"
	return *new(text.Format)
}

func (t *Table) getMaxColumnLengthForMerging(colIdx int) int { _ = "STUB: not implemented"; return 0 }

// getMergedColumnIndices returns a map of colIdx values to all the other colIdx
// values (that are being merged) and their lengths.
func (t *Table) getMergedColumnIndices(row rowStr, hint renderHint) mergedColumnIndices {
	_ = "STUB: not implemented"
	return *new(mergedColumnIndices)
}

func (t *Table) getRow(rowIdx int, hint renderHint) rowStr {
	_ = "STUB: not implemented"
	return *new(rowStr)
}

func (t *Table) getRowConfig(hint renderHint) RowConfig {
	_ = "STUB: not implemented"
	return *new(RowConfig)
}

func (t *Table) getSeparatorColors(hint renderHint) text.Colors {
	_ = "STUB: not implemented"
	return *new(text.Colors)
}

func (t *Table) getVAlign(colIdx int, hint renderHint) text.VAlign {
	_ = "STUB: not implemented"
	return *new(text.VAlign)
}

func (t *Table) hasHiddenColumns() bool { _ = "STUB: not implemented"; return false }

func (t *Table) hasRowPainter() bool { _ = "STUB: not implemented"; return false }

func (t *Table) hideColumns() map[int]int { _ = "STUB: not implemented"; return nil }

// hide columns as directed

// reset numColumns to the new number of columns

func (t *Table) isIndexColumn(colIdx int, hint renderHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) render(out *strings.Builder) string { _ = "STUB: not implemented"; return "" }

func (t *Table) shouldMergeCellsHorizontallyAbove(row rowStr, colIdx int, hint renderHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) shouldMergeCellsHorizontallyBelow(row rowStr, colIdx int, hint renderHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) shouldMergeCellsVerticallyAbove(colIdx int, hint renderHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) shouldMergeCellsVerticallyBelow(colIdx int, hint renderHint) int {
	_ = "STUB: not implemented"
	return 0
}

func (t *Table) shouldSeparateRows(rowIdx int, numRows int) bool {
	_ = "STUB: not implemented"
	// not asked to separate rows and no manually added separator
	return false
}

// last row of page

// last row of table

func (t *Table) wrapRow(row rowStr) (int, rowStr) {
	_ = "STUB: not implemented"
	return 0, *new(rowStr)
}
