package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomExtraSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17H2V3h2v12h9v2Zm5.4 4L15 14H5.5V3h6v6h4.225l3.975 8.1l2.45-1.125l1.425 2.625L18.4 21Z"/>`),
		g.Group(children),
	)
}