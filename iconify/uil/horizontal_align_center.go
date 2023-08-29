package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10h-2V7a1 1 0 0 0-1-1h-5V3a1 1 0 0 0-2 0v3H6a1 1 0 0 0-1 1v3H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h8v3a1 1 0 0 0 2 0v-3h8a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1ZM7 8h10v2H7Zm13 8H4v-4h16Z"/>`),
		g.Group(children),
	)
}