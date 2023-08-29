package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldPlusLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c0 4-7 7-7 7s-7-3-7-7V5c1.5.167 5 0 7-2c2 2 5.5 2.167 7 2v9zm-7-5v3m0 3v-3m0 0h3m-3 0H9"/>`),
		g.Group(children),
	)
}