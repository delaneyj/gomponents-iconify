package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.29 16.29L13 18.59V5.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42L11 5.41v13.18l-2.29-2.3a1 1 0 1 0-1.42 1.42l4 4a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0-1.42-1.42Z"/>`),
		g.Group(children),
	)
}