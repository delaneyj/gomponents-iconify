package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 21l4-4l-4-4"/><path d="M21 17H10a3 3 0 0 1-3-3V3"/><path d="M11 7L7 3L3 7"/></g>`),
		g.Group(children),
	)
}