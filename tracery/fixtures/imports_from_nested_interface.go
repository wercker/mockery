package test

import (
	"github.com/wercker/tracery/tracery/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
