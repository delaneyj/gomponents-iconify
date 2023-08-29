package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TodayOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16.5q-1.05 0-1.775-.725T6.5 14q0-1.05.725-1.775T9 11.5q1.05 0 1.775.725T11.5 14q0 1.05-.725 1.775T9 16.5ZM3 22V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}