package example

import "github.com/wercker/tracery/tracery"

type Example interface {
	GetWalker(string) (tracery.WalkerVisitor, error)
	MustWalker(tracery.WalkerVisitor, error) tracery.WalkerVisitor
}
