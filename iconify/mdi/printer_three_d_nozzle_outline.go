package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterThreeDNozzleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10v5h2v6h-2.5L13 17h-2l-3.5-4H5V7h2V2m3 20H2v-2h8a1 1 0 0 0 1-1v-1h2v1a3 3 0 0 1-3 3M7 9v2h1.5l3.5 4l3.5-4H17V9h-2V4H9v5H7Z"/>`),
		g.Group(children),
	)
}