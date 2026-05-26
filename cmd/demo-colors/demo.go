package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	tw := table.NewWriter()
	tw.AppendRows([]table.Row{
		{renderGrid(false)},
		{renderGrid(true)},
	})
	tw.SetTitle("256-Color Palette")
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.Style().Title.Align = text.AlignCenter
	fmt.Println(tw.Render())
}

func alignCode(code int) string { _ = "STUB: not implemented"; return "" }

func blankTable() table.Writer { _ = "STUB: not implemented"; return *new(table.Writer) }

func buildRow(start, end int, isBackground bool) table.Row {
	_ = "STUB: not implemented"
	return *new(table.Row)
}

func cellValue(code int, isBackground bool) string { _ = "STUB: not implemented"; return "" }

func renderGrid(isBackground bool) string { _ = "STUB: not implemented"; return "" }

// Standard 16 colors (0-15)

// RGB cube colors (16-231) - 216 colors in 6 blocks of 36

// Grayscale colors (232-255) - 24 colors
