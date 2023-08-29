package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogicOr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12h-6M2 9h7m-7 6h7M8 5c10.667 2.1 10.667 12.6 0 14c1.806-4.667 1.806-9.333 0-14z"/>`),
		g.Group(children),
	)
}