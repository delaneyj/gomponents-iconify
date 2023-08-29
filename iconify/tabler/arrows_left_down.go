package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3L3 7l4 4"/><path d="M3 7h11a3 3 0 0 1 3 3v11"/><path d="m13 17l4 4l4-4"/></g>`),
		g.Group(children),
	)
}