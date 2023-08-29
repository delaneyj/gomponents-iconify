package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 4.75l3.5 3.5l-3.5 3.5"/><path d="M16.75 8.25h-6a4 4 0 0 0-4 4v7"/></g>`),
		g.Group(children),
	)
}