package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueuePlayNextSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2v-3h3v-2h-3V7h-2v3H8v2h3v3Zm8.5 7.5L18 21l3-3l-3-3l1.5-1.5L24 18l-4.5 4.5ZM8 21v-2H2V3h20v9h-5v7h-2v2H8Z"/>`),
		g.Group(children),
	)
}