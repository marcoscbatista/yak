// Package reader provides reader abstractions and implementations.
package reader

type Reader interface {
	Read(p string) (c string, e error)
}
