package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.975 9H.025a4 4 0 0 1 .902-2.113l1.11-1.33A2 2 0 0 0 2.5 4.275V1a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v3.276a2 2 0 0 0 .464 1.28l1.109 1.331A4 4 0 0 1 9.975 9zM10 16v2a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2h10zm-8-5h6v3H2v-3z"/>`),
		g.Group(children),
	)
}