package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 6h8v16h-8V6M2 4h20V2H2v2m0 4h10V6H2v2m7 14h3V10H9v12m-7 0h5V10H2v12Z"/>`),
		g.Group(children),
	)
}