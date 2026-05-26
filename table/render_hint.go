package table

// renderHint has hints for the Render*() logic
type renderHint struct {
	isAutoIndexColumn bool // auto-index column?
	isAutoIndexRow    bool // auto-index row?
	isBorderBottom    bool // bottom-border?
	isBorderTop       bool // top-border?
	isFirstRow        bool // first-row of header/footer/regular-rows?
	isFooterRow       bool // footer row?
	isHeaderRow       bool // header row?
	isLastLineOfRow   bool // last-line of the current row?
	isLastRow         bool // last-row of header/footer/regular-rows?
	isSeparatorRow    bool // separator row?
	isTitleRow        bool // title row?
	rowLineNumber     int  // the line number for a multi-line row
	rowNumber         int  // the row number/index
	separatorType     separatorType
}

func (h *renderHint) isBorderOrSeparator() bool { _ = "STUB: not implemented"; return false }

func (h *renderHint) isRegularRow() bool { _ = "STUB: not implemented"; return false }

func (h *renderHint) isRegularNonSeparatorRow() bool { _ = "STUB: not implemented"; return false }

func (h *renderHint) isHeaderOrFooterSeparator() bool { _ = "STUB: not implemented"; return false }

func (h *renderHint) isLastLineOfLastRow() bool { _ = "STUB: not implemented"; return false }

type renderMode string

const (
	renderModeDefault  renderMode = "default"
	renderModeCSV      renderMode = "csv"
	renderModeMarkdown renderMode = "markdown"
	renderModeTSV      renderMode = "tsv"
	renderModeHTML     renderMode = "html"
)
