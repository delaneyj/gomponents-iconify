package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderBoolDescendingVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3c-1.11 0-2 .89-2 2v4c0 1.11.89 2 2 2h4c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2m.2 1.5l1.06 1.05L5.27 9.5L2.74 6.95L3.81 5.9l1.47 1.49M4 13c-1.11 0-2 .89-2 2v4c0 1.11.89 2 2 2h4c1.11 0 2-.89 2-2v-4c0-1.11-.89-2-2-2m-4 2h4v4H4m8-14h10v2H12m0 12v-2h10v2m-10-8h10v2H12Z"/>`),
		g.Group(children),
	)
}