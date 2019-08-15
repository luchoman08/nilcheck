// Package addcheck defines an Analyzer that reports integer additions
package addcheck

import (
	"errors"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "addlint",
	Doc:  "reports integer additions",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}
