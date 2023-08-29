package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h8v2H8v-2M2 5.6v4.8c0 .9.5 1.6 1.2 1.6h17.7c.6 0 1.2-.7 1.2-1.6V5.6C22 4.7 21.5 4 20.8 4H3.2C2.5 4 2 4.7 2 5.6M10 9V7H9v2h1M5 9h2V7H5v2m15 1H4V6h16v4Z"/>`),
		g.Group(children),
	)
}