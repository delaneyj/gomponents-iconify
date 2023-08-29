package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L15.2 18H6V8.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM2 22V6h2v14h14v2H2Zm18.7-4.125L8.825 6H20V4H8v1.175L6.125 3.3V2H22v15.875h-1.3Z"/>`),
		g.Group(children),
	)
}