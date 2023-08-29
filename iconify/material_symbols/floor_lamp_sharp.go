package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19v-8H4.65l2.775-9h9.15l2.775 9H13v8h-2Zm-3 3v-2h8v2H8Z"/>`),
		g.Group(children),
	)
}