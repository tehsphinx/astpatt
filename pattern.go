package astpatt

import (
	"encoding/json"
	"fmt"

	"github.com/pmezard/go-difflib/difflib"

	"github.com/tehsphinx/astrav"
)

// ExtractPattern extracts the pattern from a given ast package
func ExtractPattern(pkg *astrav.Package) *Pattern {
	pattern := NewPattern()
	pattern.Populate(pkg)
	return pattern
}

// MatchPatterns matches a package against multiple valid patterns.
// If one of the valid patterns matches true is returned.
func MatchPatterns(valid []*Pattern, pkg *astrav.Package) bool {
	for _, pattern := range valid {
		if pattern.MatchPkg(pkg) {
			return true
		}
	}
	return false
}

// DiffPatterns matches a package against multiple valid patterns.
// If one of the valid patterns matches true is returned.
// DiffPatterns additionally returns a diff of the patterns.
func DiffPatterns(valid []*Pattern, pkg *astrav.Package) (string, bool) {
	var minDiff string
	for _, pattern := range valid {
		diff, ok := pattern.DiffPkg(pkg)
		if ok {
			return "", true
		}
		if minDiff == "" || len(diff) < len(minDiff) {
			minDiff = diff
		}
	}
	return minDiff, false
}

// NewPattern creates a new pattern
func NewPattern() *Pattern {
	return &Pattern{}
}

// Pattern implements a solutions pattern
type Pattern struct {
	parentNode

	Name string `json:"-"`
}

// MatchPkg checks the pattern against given parent node
func (s *Pattern) MatchPkg(pkg *astrav.Package) bool {
	pkgPatt := ExtractPattern(pkg)
	return s.Match(pkgPatt)
}

// DiffPkg checks the pattern against given parent node
func (s *Pattern) DiffPkg(pkg *astrav.Package) (string, bool) {
	pkgPatt := ExtractPattern(pkg)
	if s.Match(pkgPatt) {
		return "", true
	}
	return getDiff(s.String(), pkgPatt.String()), false
}

// Match checks the pattern against another pattern
func (s *Pattern) Match(node Node) bool {
	return s.Nodes.Match(node)
}

func (s *Pattern) String() string {
	b, _ := json.MarshalIndent(*s, "", "  ")
	return string(b)
}

func getDiff(expected, current string) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(expected),
		B:        difflib.SplitLines(current),
		FromFile: "Expected",
		ToFile:   "Current",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return fmt.Sprintf("error while diffing strings: %s", err)
	}
	return text
}
