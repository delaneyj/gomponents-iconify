package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h8l2 2h10v2H11.175l-2-2H4v12l2.4-8h17.1l-3 10H2Zm4.1-2H19l1.8-6H7.9l-1.8 6Zm0 0l1.8-6l-1.8 6ZM4 8V6v2Z"/>`),
		g.Group(children),
	)
}