package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 19H20a8 8 0 1 0 0 16h16a8 8 0 0 0 6-13.292"/><path d="M6 24.292A8 8 0 0 1 12 11h16a8 8 0 1 1 0 16H18"/></g>`),
		g.Group(children),
	)
}