package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterThreeDNozzleAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10v6h2v5h-2.5L13 17h-2l-3.5-4H5V8h2V2m3 20H2v-2h8c.6 0 1-.5 1-1v-1h2v1c0 1.7-1.3 3-3 3m11-9V7h2v6h-2m0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}