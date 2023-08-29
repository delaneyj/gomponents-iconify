package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCardOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v8H4v6h10v2H2ZM4 8h16V6H4v2Zm15 14v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}