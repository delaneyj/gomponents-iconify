package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6.1L19 5v11a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3c.77 0 1.47.29 2 .76V9.04l-9 1V17a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3c.77 0 1.47.29 2 .76V6.1M9 7v2.03l9-.99V6.09L9 7M8 17a2 2 0 0 0-2-2a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2m10-1a2 2 0 0 0-2-2a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}