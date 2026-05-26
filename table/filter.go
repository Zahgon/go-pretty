package table

// FilterBy defines what to filter (Column Name or Number), how to filter (Operator),
// and the value to compare against.
type FilterBy struct {
	// Name is the name of the Column as it appears in the first Header row.
	// If a Header is not provided, or the name is not found in the header, this
	// will not work.
	Name string
	// Number is the Column # from left. When specified, it overrides the Name
	// property. If you know the exact Column number, use this instead of Name.
	Number int

	// Operator defines how to compare the column value against the Value.
	Operator FilterOperator

	// Value is the value to compare against. The type should match the expected
	// comparison type (string for string operations, numeric for numeric operations).
	// For Contains, StartsWith, EndsWith, and RegexMatch, Value should be a string.
	// For numeric comparisons (Equal, NotEqual, GreaterThan, etc.), Value can be
	// a number (int, float64) or a string representation of a number.
	Value interface{}

	// IgnoreCase makes string comparisons case-insensitive (only applies to
	// string-based operators).
	IgnoreCase bool

	// CustomFilter is a function that can be used to filter rows in a custom
	// manner. Note that:
	// * This overrides and ignores the Operator, Value, and IgnoreCase settings
	// * This is called after the column contents are converted to string form
	// * This function is expected to return:
	//   * true => include the row
	//   * false => exclude the row
	//
	// Use this when the default filtering logic is not sufficient.
	CustomFilter func(cellValue string) bool
}

// FilterOperator defines how to filter.
type FilterOperator int

const (
	// Equal filters rows where the column value equals the Value.
	Equal FilterOperator = iota
	// NotEqual filters rows where the column value does not equal the Value.
	NotEqual
	// GreaterThan filters rows where the column value is greater than the Value.
	GreaterThan
	// GreaterThanOrEqual filters rows where the column value is greater than or equal to the Value.
	GreaterThanOrEqual
	// LessThan filters rows where the column value is less than the Value.
	LessThan
	// LessThanOrEqual filters rows where the column value is less than or equal to the Value.
	LessThanOrEqual
	// Contains filters rows where the column value contains the Value (string search).
	Contains
	// NotContains filters rows where the column value does not contain the Value (string search).
	NotContains
	// StartsWith filters rows where the column value starts with the Value.
	StartsWith
	// EndsWith filters rows where the column value ends with the Value.
	EndsWith
	// RegexMatch filters rows where the column value matches the Value as a regular expression.
	RegexMatch
	// RegexNotMatch filters rows where the column value does not match the Value as a regular expression.
	RegexNotMatch
)

func (t *Table) parseFilterBy(filterBy []FilterBy) []FilterBy {
	_ = "STUB: not implemented"
	return nil
}

// Parse from raw header rows

func (t *Table) matchesFiltersRaw(row Row, filters []FilterBy) bool {
	_ = "STUB: not implemented"
	// All filters must match (AND logic)
	return false
}

func (t *Table) matchesFilterRaw(row Row, filter FilterBy) bool {
	_ = "STUB: not implemented"
	return false
}

// Use custom filter if provided

// Use operator-based filtering

func (t *Table) matchesOperator(cellValue string, filter FilterBy) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) compareEqual(cellValue string, filterValue interface{}, ignoreCase bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) compareNumeric(cellValue string, filterValue interface{}, compareFunc func(float64, float64) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Try to convert to string and parse

func (t *Table) compareContains(cellValue string, filterValue interface{}, ignoreCase bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) compareStartsWith(cellValue string, filterValue interface{}, ignoreCase bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) compareEndsWith(cellValue string, filterValue interface{}, ignoreCase bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Table) compareRegexMatch(cellValue string, filterValue interface{}, ignoreCase bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Compile the regex pattern

// If regex compilation fails, fall back to simple string matching
