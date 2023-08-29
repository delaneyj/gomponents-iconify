package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multistate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="4" cy="4" r="3" fill="currentColor"/><path fill="currentColor" d="M4 15a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm8 13a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2Zm8-11a3 3 0 1 1 3-3a3 3 0 0 1-3 3Zm0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2Z"/><circle cx="4" cy="20" r="3" fill="currentColor"/><circle cx="12" cy="4" r="3" fill="currentColor"/><circle cx="12" cy="12" r="3" fill="currentColor"/><circle cx="20" cy="12" r="3" fill="currentColor"/><circle cx="20" cy="20" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}