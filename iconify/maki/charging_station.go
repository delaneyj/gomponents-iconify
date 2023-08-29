package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.646 7.801L7.752.838a.826.826 0 0 1 1.463.705L8.086 5.684A.25.25 0 0 0 8.327 6h3.42a.753.753 0 0 1 .607 1.199l-5.106 6.963a.826.826 0 0 1-1.463-.705l1.129-4.141A.25.25 0 0 0 6.673 9h-3.42a.753.753 0 0 1-.607-1.199Z"/>`),
		g.Group(children),
	)
}