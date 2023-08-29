package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Institution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19h20v3H2zM12 2L2 6v2h20V6zm5 8h3v7h-3zm-6.5 0h3v7h-3zM4 10h3v7H4z"/>`),
		g.Group(children),
	)
}