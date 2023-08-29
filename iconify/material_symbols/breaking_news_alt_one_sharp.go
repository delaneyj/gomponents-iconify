package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingNewsAltOneSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h5v-2H6v2Zm10 0h2v-2h-2v2ZM6 13h5v-2H6v2Zm10 0h2V7h-2v6ZM6 9h5V7H6v2ZM2 21V3h20v18H2Z"/>`),
		g.Group(children),
	)
}