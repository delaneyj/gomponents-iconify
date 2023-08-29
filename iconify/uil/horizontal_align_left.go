package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10h-5V7a1 1 0 0 0-1-1H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-3h17a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1ZM4 8h10v2H4Zm16 8H4v-4h16Z"/>`),
		g.Group(children),
	)
}