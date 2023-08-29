package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 8l2-2M8 8l2 2M8 8h6a2 2 0 0 1 2 2v1m0 5l-2 2m2-2l-2-2m2 2h-6a2 2 0 0 1-2-2v-1"/><circle r="10" transform="matrix(-1 0 0 1 12 12)"/></g>`),
		g.Group(children),
	)
}