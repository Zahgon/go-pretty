package table

// BoxStyle defines the characters/strings to use to render the borders and
// separators for the Table.
type BoxStyle struct {
	BottomLeft       string
	BottomRight      string
	BottomSeparator  string
	EmptySeparator   string
	Left             string
	LeftSeparator    string
	MiddleHorizontal string
	MiddleSeparator  string
	MiddleVertical   string
	PaddingLeft      string
	PaddingRight     string
	PageSeparator    string
	Right            string
	RightSeparator   string
	TopLeft          string
	TopRight         string
	TopSeparator     string
	UnfinishedRow    string

	// Horizontal lets you customize the horizontal lines for the Table
	// in a more granular way than the MiddleHorizontal string. Setting
	// this to a non-nil value will override MiddleHorizontal.
	Horizontal *BoxStyleHorizontal
}

// BoxStyleHorizontal defines the characters/strings to use to render the
// horizontal lines for the Table.
type BoxStyleHorizontal struct {
	TitleTop     string
	TitleBottom  string // overrides HeaderTop/RowTop
	HeaderTop    string
	HeaderMiddle string
	HeaderBottom string // overrides RowTop
	RowTop       string
	RowMiddle    string
	RowBottom    string
	FooterTop    string // overrides RowBottom
	FooterMiddle string
	FooterBottom string
}

// NewBoxStyleHorizontal creates a new BoxStyleHorizontal with the given
// horizontal string.
func NewBoxStyleHorizontal(horizontal string) *BoxStyleHorizontal {
	_ = "STUB: not implemented"
	return nil
}

var (
	// StyleBoxDefault defines a Boxed-Table like below:
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   # | FIRST NAME | LAST NAME | SALARY |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |   1 | Arya       | Stark     |   3000 |                             |
	//  |  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
	//  | 300 | Tyrion     | Lannister |   5000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	//  |     |            | TOTAL     |  10000 |                             |
	//  +-----+------------+-----------+--------+-----------------------------+
	StyleBoxDefault = BoxStyle{
		BottomLeft:       "+",
		BottomRight:      "+",
		BottomSeparator:  "+",
		EmptySeparator:   " ",
		Left:             "|",
		LeftSeparator:    "+",
		MiddleHorizontal: "-",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "|",
		RightSeparator:   "+",
		TopLeft:          "+",
		TopRight:         "+",
		TopSeparator:     "+",
		UnfinishedRow:    " ~",
	}

	// StyleBoxBold defines a Boxed-Table like below:
	//  ┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	//  ┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃
	//  ┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃
	//  ┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃
	//  ┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
	//  ┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃
	//  ┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
	StyleBoxBold = BoxStyle{
		BottomLeft:       "┗",
		BottomRight:      "┛",
		BottomSeparator:  "┻",
		EmptySeparator:   " ",
		Left:             "┃",
		LeftSeparator:    "┣",
		MiddleHorizontal: "━",
		MiddleSeparator:  "╋",
		MiddleVertical:   "┃",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "┃",
		RightSeparator:   "┫",
		TopLeft:          "┏",
		TopRight:         "┓",
		TopSeparator:     "┳",
		UnfinishedRow:    " ≈",
	}

	// StyleBoxDouble defines a Boxed-Table like below:
	//  ╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗
	//  ║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║
	//  ║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║
	//  ║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║
	//  ╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣
	//  ║     ║            ║ TOTAL     ║  10000 ║                             ║
	//  ╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝
	StyleBoxDouble = BoxStyle{
		BottomLeft:       "╚",
		BottomRight:      "╝",
		BottomSeparator:  "╩",
		EmptySeparator:   " ",
		Left:             "║",
		LeftSeparator:    "╠",
		MiddleHorizontal: "═",
		MiddleSeparator:  "╬",
		MiddleVertical:   "║",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "║",
		RightSeparator:   "╣",
		TopLeft:          "╔",
		TopRight:         "╗",
		TopSeparator:     "╦",
		UnfinishedRow:    " ≈",
	}

	// StyleBoxLight defines a Boxed-Table like below:
	//  ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  └─────┴────────────┴───────────┴────────┴─────────────────────────────┘
	StyleBoxLight = BoxStyle{
		BottomLeft:       "└",
		BottomRight:      "┘",
		BottomSeparator:  "┴",
		EmptySeparator:   " ",
		Left:             "│",
		LeftSeparator:    "├",
		MiddleHorizontal: "─",
		MiddleSeparator:  "┼",
		MiddleVertical:   "│",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "│",
		RightSeparator:   "┤",
		TopLeft:          "┌",
		TopRight:         "┐",
		TopSeparator:     "┬",
		UnfinishedRow:    " ≈",
	}

	// StyleBoxRounded defines a Boxed-Table like below:
	//  ╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮
	//  │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │   1 │ Arya       │ Stark     │   3000 │                             │
	//  │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
	//  │ 300 │ Tyrion     │ Lannister │   5000 │                             │
	//  ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
	//  │     │            │ TOTAL     │  10000 │                             │
	//  ╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯
	StyleBoxRounded = BoxStyle{
		BottomLeft:       "╰",
		BottomRight:      "╯",
		BottomSeparator:  "┴",
		EmptySeparator:   " ",
		Left:             "│",
		LeftSeparator:    "├",
		MiddleHorizontal: "─",
		MiddleSeparator:  "┼",
		MiddleVertical:   "│",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            "│",
		RightSeparator:   "┤",
		TopLeft:          "╭",
		TopRight:         "╮",
		TopSeparator:     "┬",
		UnfinishedRow:    " ≈",
	}

	// styleBoxTest defines a Boxed-Table like below:
	//  (-----^------------^-----------^--------^-----------------------------)
	//  [<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
	//  [< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
	//  [<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
	//  {-----+------------+-----------+--------+-----------------------------}
	//  [<   >|<          >|<TOTAL    >|< 10000>|<                           >]
	//  \-----v------------v-----------v--------v-----------------------------/
	styleBoxTest = BoxStyle{
		BottomLeft:       "\\",
		BottomRight:      "/",
		BottomSeparator:  "v",
		EmptySeparator:   " ",
		Left:             "[",
		LeftSeparator:    "{",
		MiddleHorizontal: "--",
		MiddleSeparator:  "+",
		MiddleVertical:   "|",
		PaddingLeft:      "<",
		PaddingRight:     ">",
		PageSeparator:    "\n",
		Right:            "]",
		RightSeparator:   "}",
		TopLeft:          "(",
		TopRight:         ")",
		TopSeparator:     "^",
		UnfinishedRow:    " ~~~",
	}
)

type separatorType int

const (
	separatorTypeTitleTop separatorType = iota
	separatorTypeTitleBottom
	separatorTypeHeaderTop
	separatorTypeHeaderMiddle
	separatorTypeHeaderBottom
	separatorTypeRowTop
	separatorTypeRowMiddle
	separatorTypeRowBottom
	separatorTypeFooterTop
	separatorTypeFooterMiddle
	separatorTypeFooterBottom
	separatorTypeCount // this should be the last value
)

func (bs *BoxStyle) ensureHorizontalInitialized() { _ = "STUB: not implemented"; return }

func (bs *BoxStyle) middleHorizontal(st separatorType) string { _ = "STUB: not implemented"; return "" }
