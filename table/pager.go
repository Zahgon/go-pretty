package table

import (
	"io"
)

// Pager lets you interact with the table rendering in a paged manner.
type Pager interface {
	// GoTo moves to the given 1-indexed page number.
	GoTo(pageNum int) string
	// Location returns the current page number in 1-indexed form.
	Location() int
	// Next moves to the next available page and returns the same.
	Next() string
	// Prev moves to the previous available page and returns the same.
	Prev() string
	// Render returns the current page.
	Render() string
	// SetOutputMirror sets up the writer to which Render() will write the
	// output other than returning.
	SetOutputMirror(mirror io.Writer)
}

type pager struct {
	index        int // 0-indexed
	pages        []string
	outputMirror io.Writer
	size         int
}

func (p *pager) GoTo(pageNum int) string { _ = "STUB: not implemented"; return "" }

func (p *pager) Location() int { _ = "STUB: not implemented"; return 0 }

func (p *pager) Next() string { _ = "STUB: not implemented"; return "" }

func (p *pager) Prev() string { _ = "STUB: not implemented"; return "" }

func (p *pager) Render() string { _ = "STUB: not implemented"; return "" }

func (p *pager) SetOutputMirror(mirror io.Writer) { _ = "STUB: not implemented"; return }
