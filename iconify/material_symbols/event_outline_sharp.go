package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 18q-1.05 0-1.775-.725T12 15.5q0-1.05.725-1.775T14.5 13q1.05 0 1.775.725T17 15.5q0 1.05-.725 1.775T14.5 18ZM3 22V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}