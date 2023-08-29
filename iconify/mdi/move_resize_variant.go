package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveResizeVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.88.46L.46 1.88L5.59 7H2v2h7V2H7v3.59M11 7v2h10v6h2V9a2 2 0 0 0-2-2M7 11v10a2 2 0 0 0 2 2h6v-2H9V11m6.88 3.46l-1.42 1.42L19.6 21H17v2h6v-6h-2v2.59"/>`),
		g.Group(children),
	)
}