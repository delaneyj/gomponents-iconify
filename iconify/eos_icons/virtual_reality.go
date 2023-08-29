package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualReality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10h2v4h-2zM1 10h2v4H1zm17-1a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2a2 2 0 0 0 2 2h1.028l1.602-2.776a1.582 1.582 0 0 1 2.74 0L14.972 17H16a2 2 0 0 0 2-2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2Zm-1 3H7v-2h10Z"/>`),
		g.Group(children),
	)
}