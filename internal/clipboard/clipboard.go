// Package clipboard provides clipboard abstractions and implementations.
package clipboard

import "os"

type Clipboard interface {
	Init() error
	Write(text string) error
}

func NewClipboard() (Clipboard, error) {
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		cb := &WaylandClipboard{}
		return cb, cb.Init()
	}
	cb := &GoDesignClipboard{}
	return cb, cb.Init()
}
