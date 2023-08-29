package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trekking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0M7 21l2-4m4 4v-4l-3-3l1-6l3 4l3 2"/><path d="m10 14l-1.827-1.218a2 2 0 0 1-.831-2.15l.28-1.117A2 2 0 0 1 9.561 8H11l4 1l3-2m-1 5v9m-1-1h2"/></g>`),
		g.Group(children),
	)
}