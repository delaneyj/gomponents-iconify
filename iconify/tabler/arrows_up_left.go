package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 7l-4-4l-4 4"/><path d="M17 3v11a3 3 0 0 1-3 3H3"/><path d="m7 13l-4 4l4 4"/></g>`),
		g.Group(children),
	)
}