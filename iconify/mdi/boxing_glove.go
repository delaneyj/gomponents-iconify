package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxingGlove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16V6h3v10h-3M12 4H7S2 4 2 8v6c0 1.77 1 2.76 2.07 3.31A3.996 3.996 0 0 1 8 14h3v2H8a2 2 0 0 0-2 2a2 2 0 0 0 2 2h5c4 0 4-4 4-4V6s-1-2-5-2Z"/>`),
		g.Group(children),
	)
}