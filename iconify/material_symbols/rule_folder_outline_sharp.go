package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleFolderOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h8l2 2h10v14H2Zm2-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Zm3.825-1.625l4.95-4.95L11.35 10l-3.525 3.55L6.4 12.125L5 13.55l2.825 2.825Zm6.575 0l1.6-1.6l1.6 1.6l1.4-1.4l-1.6-1.6l1.6-1.6l-1.4-1.4l-1.6 1.6l-1.6-1.6l-1.4 1.4l1.6 1.6l-1.6 1.6l1.4 1.4Z"/>`),
		g.Group(children),
	)
}