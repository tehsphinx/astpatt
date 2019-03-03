package astpatt

import (
	"errors"
	"testing"

	"github.com/tehsphinx/astrav"
)

func TestPattern_Match(t *testing.T) {
	pkg, err := getPackage("solutions/1")
	if err != nil {
		t.Error(err)
	}

	pattern := ExtractPattern(pkg)
	if !pattern.MatchPkg(pkg) {
		t.Fail()
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
