package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Functions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16v2H7.14l8.275 7l-8.275 7H20v2H4v-1.964L12.318 12L4 4.964V3Z"/>`),
		g.Group(children),
	)
}