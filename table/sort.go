package table

// SortBy defines What to sort (Column Name or Number), and How to sort (Mode).
type SortBy struct {
	// Name is the name of the Column as it appears in the first Header row.
	// If a Header is not provided, or the name is not found in the header, this
	// will not work.
	Name string
	// Number is the Column # from left. When specified, it overrides the Name
	// property. If you know the exact Column number, use this instead of Name.
	Number int

	// Mode tells the Writer how to Sort. Asc/Dsc/etc.
	Mode SortMode

	// IgnoreCase makes sorting case-insensitive
	IgnoreCase bool

	// CustomLess is a function that can be used to sort the column in a custom
	// manner. Note that:
	// * This overrides and ignores the Mode and IgnoreCase settings
	// * This is called after the column contents are converted to string form
	// * This function is expected to return:
	//   * -1 => when iStr comes before jStr
	//   *  0 => when iStr and jStr are considered equal
	//   *  1 => when iStr comes after jStr
	//
	// Use this when the default sorting logic is not sufficient.
	CustomLess func(iStr string, jStr string) int
}

// SortMode defines How to sort.
type SortMode int

const (
	// Asc sorts the column in Ascending order alphabetically.
	Asc SortMode = iota
	// AscAlphaNumeric sorts the column in Ascending order alphabetically and
	// then numerically.
	AscAlphaNumeric
	// AscNumeric sorts the column in Ascending order numerically.
	AscNumeric
	// AscNumericAlpha sorts the column in Ascending order numerically and
	// then alphabetically.
	AscNumericAlpha
	// Dsc sorts the column in Descending order alphabetically.
	Dsc
	// DscAlphaNumeric sorts the column in Descending order alphabetically and
	// then numerically.
	DscAlphaNumeric
	// DscNumeric sorts the column in Descending order numerically.
	DscNumeric
	// DscNumericAlpha sorts the column in Descending order numerically and
	// then alphabetically.
	DscNumericAlpha
)

// getSortedRowIndices sorts and returns the row indices in Sorted order as
// directed by Table.sortBy which can be set using Table.SortBy(...)
func (t *Table) getSortedRowIndices() []int { _ = "STUB: not implemented"; return nil }

// extract the values/cells from the rows for comparison

// compare and choose whether to continue

// if the values are not equal, return the result immediately

// if the values are equal, continue to the next column

func (t *Table) parseSortBy(sortBy []SortBy) []SortBy { _ = "STUB: not implemented"; return nil }

func less(iVal string, jVal string, sb SortBy) (bool, bool) {
	_ = "STUB: not implemented"
	return false,

		// use the custom less function to compare the values
		false
}

// rc == 0

// if the values are equal, return fast to continue to next column

// otherwise, use the default sorting logic defined by Mode and IgnoreCase

// AscAlphaNumeric, AscNumericAlpha, DscAlphaNumeric, DscNumericAlpha

func lessAlphabetic(iVal string, jVal string, sb SortBy) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// when two strings are case-insensitive identical, compare them casesensitive.
// That makes sure to get a consistent sorting

// Dsc, DscAlphaNumeric, DscNumericAlpha

// Dsc, DscAlphaNumeric, DscNumericAlpha

func lessMixedMode(iVal string, jVal string, sb SortBy) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// both are alphanumeric

// iVal == "abc"; jVal == 5

// AscNumericAlpha, DscNumericAlpha

// iVal == 5; jVal	== "abc"

// AscNumericAlpha, DscNumericAlpha:

// both values numeric

func lessNumeric(iVal string, jVal string, sb SortBy) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func lessNumericVal(iVal float64, jVal float64, sb SortBy) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// DscNumeric, DscAlphaNumeric, DscNumericAlpha
