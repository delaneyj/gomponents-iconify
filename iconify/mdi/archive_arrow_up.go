package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21h16V8H4m10 7v3h-4v-3H7l5-5l5 5M3 3h18v4H3"/>`),
		g.Group(children),
	)
}