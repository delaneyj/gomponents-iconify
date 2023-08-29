package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RestTimeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6v2a6 6 0 1 0 5.996 6.225L17 14h2a8 8 0 1 1-16 0c0-4.335 3.58-8 8-8Zm10-4v2l-5.327 6H21v2h-8v-2l5.326-6H13V2h8Z"/>`),
		g.Group(children),
	)
}