package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueuePlayNextOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 22.5L18 21l3-3l-3-3l1.5-1.5L24 18l-4.5 4.5ZM8 21v-2H2V3h20v9h-2V5H4v12h13v2h-2v2H8Zm3-6h2v-3h3v-2h-3V7h-2v3H8v2h3v3Zm-7 2V5v12Z"/>`),
		g.Group(children),
	)
}