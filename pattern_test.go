package astpatt

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/astrav"
)

type test struct {
	solution string
	match    bool
}

var matchTests = []struct {
	patterns []string
	tests    []test
}{
	{
		patterns: []string{"solutions/twofer/patterns/1", "solutions/twofer/patterns/2"},
		tests: []test{
			{solution: "solutions/twofer/1", match: true},
			{solution: "solutions/twofer/2", match: false},
			{solution: "solutions/twofer/3", match: true},
			{solution: "solutions/twofer/4", match: true},
			{solution: "solutions/twofer/5", match: false},
			{solution: "solutions/twofer/6", match: false},
			{solution: "solutions/twofer/7", match: false},
			{solution: "solutions/twofer/8", match: true},
			{solution: "solutions/twofer/9", match: true},
			{solution: "solutions/twofer/10", match: true},
			{solution: "solutions/twofer/11", match: true},
			{solution: "solutions/twofer/12", match: true},
		},
	},
	// {
	// 	patterns: []string{"solutions/hamming/patterns/1", "solutions/hamming/patterns/2"},
	// 	tests: []test{
	// 		{solution: "solutions/hamming/1", match: true},
	// 		{solution: "solutions/hamming/2", match: true},
	// 		{solution: "solutions/hamming/3", match: true},
	// 		{solution: "solutions/hamming/4", match: false},
	// 		{solution: "solutions/hamming/5", match: false},
	// 		{solution: "solutions/hamming/6", match: true},
	// 		{solution: "solutions/hamming/7", match: true},
	// 		{solution: "solutions/hamming/8", match: false},
	// 		{solution: "solutions/hamming/9", match: true},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/raindrops/2", "solutions/raindrops/4"},
	// 	tests: []test{
	// 		{solution: "solutions/raindrops/1", match: false},
	// 		{solution: "solutions/raindrops/2", match: true},
	// 		{solution: "solutions/raindrops/3", match: false},
	// 		{solution: "solutions/raindrops/4", match: true},
	// 		{solution: "solutions/raindrops/5", match: true},
	// 		{solution: "solutions/raindrops/6", match: false},
	// 		{solution: "solutions/raindrops/7", match: false},
	// 		{solution: "solutions/raindrops/8", match: false},
	// 		{solution: "solutions/raindrops/9", match: false},
	// 		{solution: "solutions/raindrops/10", match: false},
	// 		{solution: "solutions/raindrops/11", match: false},
	// 		{solution: "solutions/raindrops/12", match: false},
	// 		{solution: "solutions/raindrops/13", match: true},
	// 		{solution: "solutions/raindrops/14", match: false},
	// 		{solution: "solutions/raindrops/15", match: false},
	// 		{solution: "solutions/raindrops/16", match: false},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/raindrops/7"},
	// 	tests: []test{
	// 		{solution: "solutions/raindrops/1", match: false},
	// 		{solution: "solutions/raindrops/2", match: false},
	// 		{solution: "solutions/raindrops/3", match: false},
	// 		{solution: "solutions/raindrops/4", match: false},
	// 		{solution: "solutions/raindrops/5", match: false},
	// 		{solution: "solutions/raindrops/6", match: false},
	// 		{solution: "solutions/raindrops/7", match: true},
	// 		{solution: "solutions/raindrops/8", match: false},
	// 		{solution: "solutions/raindrops/9", match: false},
	// 		{solution: "solutions/raindrops/10", match: false},
	// 		{solution: "solutions/raindrops/11", match: false},
	// 		{solution: "solutions/raindrops/12", match: false},
	// 		{solution: "solutions/raindrops/13", match: false},
	// 		{solution: "solutions/raindrops/14", match: false},
	// 		{solution: "solutions/raindrops/15", match: false},
	// 		{solution: "solutions/raindrops/16", match: false},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/raindrops/6"},
	// 	tests: []test{
	// 		{solution: "solutions/raindrops/1", match: false},
	// 		{solution: "solutions/raindrops/2", match: false},
	// 		{solution: "solutions/raindrops/3", match: false},
	// 		{solution: "solutions/raindrops/4", match: false},
	// 		{solution: "solutions/raindrops/5", match: false},
	// 		{solution: "solutions/raindrops/6", match: true},
	// 		{solution: "solutions/raindrops/7", match: false},
	// 		{solution: "solutions/raindrops/8", match: true},
	// 		{solution: "solutions/raindrops/9", match: false},
	// 		{solution: "solutions/raindrops/10", match: true},
	// 		{solution: "solutions/raindrops/11", match: true},
	// 		{solution: "solutions/raindrops/12", match: false},
	// 		{solution: "solutions/raindrops/13", match: false},
	// 		{solution: "solutions/raindrops/14", match: false},
	// 		{solution: "solutions/raindrops/15", match: false},
	// 		{solution: "solutions/raindrops/16", match: false},
	// 	},
	// },
	// {
	// 	patterns: []string{"solutions/isogram/9", "solutions/isogram/1"},
	// 	tests: []test{
	// 		{solution: "solutions/isogram/1", match: true},
	// 		{solution: "solutions/isogram/2", match: false},
	// 		{solution: "solutions/isogram/3", match: true},
	// 		{solution: "solutions/isogram/4", match: false},
	// 		{solution: "solutions/isogram/5", match: true},
	// 		{solution: "solutions/isogram/6", match: true},
	// 		{solution: "solutions/isogram/7", match: false},
	// 		{solution: "solutions/isogram/8", match: false},
	// 		{solution: "solutions/isogram/9", match: true},
	// 	},
	// },
}

func TestPattern_Match(t *testing.T) {
	for _, group := range matchTests {
		var valid []*Pattern
		for _, folder := range group.patterns {
			pkg, err := getPackage(folder)
			if err != nil {
				t.Error(err)
			}
			pattern := ExtractPattern(pkg)
			pattern.Name = folder
			valid = append(valid, pattern)
		}

		for _, test := range group.tests {
			pkg, err := getPackage(test.solution)
			if err != nil {
				t.Error(err)
			}

			diff, _, ok := DiffPatterns(valid, pkg)
			assert.Equal(t, test.match, ok, fmt.Sprintf("solution failed: %s\n%s", test.solution, diff))
		}
	}
}

func getPackage(path string) (*astrav.Package, error) {
	folder := astrav.NewFolder(path)
	packages, err := folder.ParseFolder()
	if err != nil {
		return nil, err
	}
	for _, pkg := range packages {
		return pkg, nil
	}

	return nil, errors.New("no go package found")
}
