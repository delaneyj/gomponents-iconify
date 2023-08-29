package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenLeftOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h8v18H3Zm10 0V3h8v18h-8Zm6-16h-4v14h4V5Zm-4 14h4h-4Z"/>`),
		g.Group(children),
	)
}