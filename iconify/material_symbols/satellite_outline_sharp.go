package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SatelliteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h12l-3.75-5l-3 4L9 13l-3 4Zm0-5q2.5 0 4.25-1.75T12 6h-1.7q0 1.8-1.25 3.05T6 10.3V12Zm0-3.4q1.075 0 1.813-.763T8.55 6H6v2.6ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}