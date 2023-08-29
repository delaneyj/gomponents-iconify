package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V5h14v18l-7-3l-7 3Zm2-3.05l5-2.15l5 2.15V7H5v12.95ZM19 20V3H6V1h15v19h-2ZM5 7h10H5Z"/>`),
		g.Group(children),
	)
}