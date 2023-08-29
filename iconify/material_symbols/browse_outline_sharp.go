package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h8v6H3Zm10 0V11h8v10h-8Zm-4-4Zm6-4ZM3 13V3h8v10H3Zm6-2Zm4-2V3h8v6h-8Zm2-2ZM5 19h4v-2H5v2Zm10 0h4v-6h-4v6ZM5 11h4V5H5v6Zm10-4h4V5h-4v2Z"/>`),
		g.Group(children),
	)
}