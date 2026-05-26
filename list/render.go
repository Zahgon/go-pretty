package list

import (
	"strings"
)

// Render renders the List in a human-readable "pretty" format. Example:
// | * Game Of Thrones
// |   * Winter
// |   * Is
// |   * Coming
// |     * This
// |     * Is
// |     * Known
// | * The Dark Tower
// |   * The Gunslinger
func (l *List) Render() string { _ = "STUB: not implemented"; return "" }

func (l *List) renderItem(out *strings.Builder, idx int, item *listItem, hint renderHint) {
	_ = "STUB: not implemented"
	// when working on item number 2 or more, render a newline first
	return
}

// format item.Text as directed in l.style

// convert newlines if newlines are not "\n" in l.style

// render the item.Text line by line

// render the prefix or the leading text before the actual item

// render the actual item

func (l *List) renderItemBullet(out *strings.Builder, lineIdx int, hint renderHint) {
	_ = "STUB: not implemented"

	// multi-line item.Text
	return
}

func (l *List) renderItemBulletSingleLine(out *strings.Builder, hint renderHint) {
	_ = "STUB: not implemented"
	// single-line item.Text (or first line of a multi-line item.Text)
	return
}

func (l *List) renderItemBulletPrefix(out *strings.Builder, itemIdx int, itemLevel int) {
	_ = "STUB: not implemented"
	// write a prefix if one has been set in l.style
	return
}

// render spaces and connectors until the item's position
