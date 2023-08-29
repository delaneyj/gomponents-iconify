package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddAdOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM1 21V3h18v7h-2V8H3v11h13v2H1ZM3 6h14V5H3v1Zm0 0V5v1Z"/>`),
		g.Group(children),
	)
}