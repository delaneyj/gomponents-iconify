package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MimoDisconnectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L15.2 18H17l1 1v2H6v-2l1-1H2V3.175h1.175v2.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4Zm.2-5.45L5.85 3H22v14.85h-1.3Z"/>`),
		g.Group(children),
	)
}