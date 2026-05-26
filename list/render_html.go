package list

import (
	"strings"
)

// RenderHTML renders the List in the HTML format. Example:
//
//	<ul class="go-pretty-table">
//	  <li>Game Of Thrones</li>
//	  <ul class="go-pretty-table-1">
//	    <li>Winter</li>
//	    <li>Is</li>
//	    <li>Coming</li>
//	    <ul class="go-pretty-table-2">
//	      <li>This</li>
//	      <li>Is</li>
//	      <li>Known</li>
//	    </ul>
//	  </ul>
//	  <li>The Dark Tower</li>
//	  <ul class="go-pretty-table-1">
//	    <li>The Gunslinger</li>
//	  </ul>
//	</ul>
func (l *List) RenderHTML() string { _ = "STUB: not implemented"; return "" }

func (l *List) htmlRenderRecursively(out *strings.Builder, idx int, item *listItem) int {
	_ = "STUB: not implemented"
	return 0
}

// indent

// un-indent
