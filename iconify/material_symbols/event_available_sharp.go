package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventAvailableSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.95 18.35L7.4 14.8l1.45-1.45l2.1 2.1l4.2-4.2l1.45 1.45l-5.65 5.65ZM3 22V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10Z"/>`),
		g.Group(children),
	)
}