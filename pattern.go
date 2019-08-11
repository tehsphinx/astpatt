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

// ExtractPatternPermutations extracts a list of patterns with nodes order exchanged where allowed.
func ExtractPatternPermutations(pkg *astrav.Package) []*Pattern {
	perms := permutations(pkg)

	var patterns []*Pattern
	for _, node := range perms {
		patterns = append(patterns, node.(*Pattern))
	}
	return patterns
}

// MatchPatterns matches a package against multiple valid patterns.
// If one of the valid patterns matches true is returned.
func MatchPatterns(valid []*Pattern, pkg *astrav.Package) bool {
	if pkg == nil {
		return false
	}
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
func DiffPatterns(valid []*Pattern, pkg *astrav.Package) (string, float64, bool) {
	var (
		minDiff  string
		maxRatio float64
	)
	for _, pattern := range valid {
		diff, ratio, ok := pattern.DiffPkg(pkg)
		if ok {
			return "", ratio, true
		}
		if minDiff == "" || maxRatio < ratio {
			minDiff = diff
			maxRatio = ratio
		}
	}
	return minDiff, maxRatio, false
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

// MatchPkg checks the pattern against given package
func (s *Pattern) MatchPkg(pkg *astrav.Package) bool {
	pkgPatt := ExtractPattern(pkg)
	return s.Match(pkgPatt)
}

// DiffPkg checks the pattern against given package
func (s *Pattern) DiffPkg(pkg *astrav.Package) (string, float64, bool) {
	pkgPatt := ExtractPattern(pkg)
	return s.DiffPattern(pkgPatt)
}

// DiffPattern checks the pattern against another pattern
func (s *Pattern) DiffPattern(pattern *Pattern) (string, float64, bool) {
	if s.Match(pattern) {
		return "", 1, true
	}
	diff, ratio := getDiff(s.String(), pattern.String())
	return diff, ratio, false
}

// Match checks the pattern against another pattern
func (s *Pattern) Match(node Node) bool {
	return s.Nodes.Match(node)
}

func (s *Pattern) String() string {
	b, _ := json.MarshalIndent(*s, "", "  ")
	return string(b)
}

func getDiff(expected, current string) (string, float64) {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(expected),
		B:        difflib.SplitLines(current),
		FromFile: "Expected",
		ToFile:   "Current",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return fmt.Sprintf("error while diffing strings: %s", err), 0
	}
	matcher := difflib.NewMatcherWithJunk(difflib.SplitLines(expected), difflib.SplitLines(current),
		false, func(s string) bool { return false })
	return text, matcher.Ratio()
}
