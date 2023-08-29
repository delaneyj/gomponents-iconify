package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13h6a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-3l-1-1h-2a1 1 0 0 0-1 1v1H8V7h2a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H7L6 1H4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2v13h7v1a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3l-1-1h-2a1 1 0 0 0-1 1v1H8v-7h5v1a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}