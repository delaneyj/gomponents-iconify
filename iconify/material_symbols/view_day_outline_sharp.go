package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDayOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-2h18v2H3Zm0-4V8h18v8H3Zm2-2h14v-4H5v4ZM3 6V4h18v2H3Zm2 8v-4v4Z"/>`),
		g.Group(children),
	)
}