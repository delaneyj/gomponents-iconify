package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-8.2h1V5h18v5.8h1V19h-2v-2H4v2H2Zm11-9h6V7h-6v3Zm-8 0h6V7H5v3Z"/>`),
		g.Group(children),
	)
}