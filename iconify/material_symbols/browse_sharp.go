package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h8v6H3Zm10 0V11h8v10h-8ZM3 13V3h8v10H3Zm10-4V3h8v6h-8Z"/>`),
		g.Group(children),
	)
}