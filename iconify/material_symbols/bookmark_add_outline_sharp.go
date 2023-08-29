package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAddOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18l-7 3V3h8v2H7v12.95l5-2.15l5 2.15V11h2v10l-7-3ZM7 5h6h-6Zm10 4V7h-2V5h2V3h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}