package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShippingPallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20h2v-2h6v2h2v-2h6v2h2v-5h-2v1h-2v-1h-2v1h-2v-1h-2v1H9v-1H7v1H5v-1H3m2-2h14V4H5Z"/>`),
		g.Group(children),
	)
}