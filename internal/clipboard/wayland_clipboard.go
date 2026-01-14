package clipboard

import (
	"os/exec"
	"strings"
)

type WaylandClipboard struct{}

func (w *WaylandClipboard) Init() error {
	_, err := exec.LookPath("wl-copy")
	return err
}

func (w *WaylandClipboard) Write(text string) error {
	cmd := exec.Command("wl-copy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
