package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MimoDisconnectOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2l1-1H2V3.175h1.175L5 5H4v11h9.2L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-5.3-5.3H17l1 1v2H6Zm14.7-3.15L18.85 16H20V5H7.85l-2-2H22v14.85h-1.3Zm-7.35-7.35Zm-4.75.9Z"/>`),
		g.Group(children),
	)
}