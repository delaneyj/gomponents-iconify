package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyPottedTreeOneTreePlantPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="4" r="3.5"/><path d="M7 5v4.5m2 4H5l-.5-4h5l-.5 4z"/></g>`),
		g.Group(children),
	)
}