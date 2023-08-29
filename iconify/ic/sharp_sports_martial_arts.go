package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSportsMartialArts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 2l-8.2 6.7l-1.21-1.04l3.6-2.08L9.41 1L8 2.41l2.74 2.74L5 8.46l-1.19 4.29L6.27 17L8 16l-2.03-3.52l.35-1.3L9.5 13l.5 9h2l.5-10L21 3.4z"/><circle cx="5" cy="5" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}