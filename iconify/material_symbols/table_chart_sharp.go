package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8V3h19v5H3Zm5 2v11H3V10h5Zm14 11h-5V10h5v11Zm-7-11v11h-5V10h5Z"/>`),
		g.Group(children),
	)
}