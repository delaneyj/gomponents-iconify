package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkRemoveOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7h-6V5h6v2Zm-9 11l-7 3V3h8v2H7v12.95l5-2.15l5 2.15V11h2v10l-7-3ZM7 5h6h-6Z"/>`),
		g.Group(children),
	)
}