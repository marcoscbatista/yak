package clipboard

import (
	"golang.design/x/clipboard"
)

type GoDesignClipboard struct{}

func (c *GoDesignClipboard) Init() error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	return nil
}

func (c *GoDesignClipboard) Write(text string) error {
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}
