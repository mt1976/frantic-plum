package commonConfig

import "strings"

func (s *Settings) GetHistory_MaxHistoryEntries() int {
	return s.History.MaxEntries
}

// GetDisplayDelimiter returns the delimiter used to separate the display elements
// Deprecated: Use GetDefault_Delimiter instead
func (s *Settings) GetDisplayDelimiter() string {
	return s.GetDefault_Delimiter()
}

// GetDefault_Delimiter returns the delimiter used to separate the display elements
func (s *Settings) GetDefault_Delimiter() string {
	if s.Display.Delimiter == "" {
		return "⋮"
	}
	return s.Display.Delimiter
}

// SEP returns the delimiter used to separate the display elements
// Deprecated: Use Delimiter instead
func (s *Settings) SEP() string {
	return s.GetDefault_Delimiter()
}

// Delimiter returns the delimiter used to separate the display elements
// Deprecated: Use GetDefault_Delimiter instead
func (s *Settings) Delimiter() string {
	return s.GetDefault_Delimiter()
}

func isTrueFalse(s string) bool {
	// We only disable the logging if the value is "true"/"t" or "yes"/"y"
	logTrue := "true"
	if strings.EqualFold(s[:1], "y") {
		logTrue = "yes"
	}

	if strings.EqualFold(s, logTrue[:1]) {
		return true
	}
	if strings.EqualFold(s, logTrue) {
		return true
	}
	return false
}
