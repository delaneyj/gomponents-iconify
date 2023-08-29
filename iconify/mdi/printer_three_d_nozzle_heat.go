package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterThreeDNozzleHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h10v5h2v6h-2.5L10 17H8l-3.5-4H2V7h2V2m19 12.5l-1.4 2.2l1.4 2.2l-2 3.1l-1.8-.9l1.5-2.2l-1.5-2.2l2-3.1l1.8.9m-4.3 0l-1.5 2.2l1.5 2.2l-2 3.1l-1.8-.9l1.4-2.2l-1.4-2.2l2-3.1l1.8.9"/>`),
		g.Group(children),
	)
}