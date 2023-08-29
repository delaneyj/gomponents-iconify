package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxingGloves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 16V6h3v10h-3zM12 4H7S2 4 2 8v6c0 1.772.979 2.757 2.07 3.306A4 4 0 0 1 8 14h3v2H8a2 2 0 1 0 0 4h5c4 0 4-4 4-4V6s-1-2-5-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}