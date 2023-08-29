package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.25 19.25l3.5-3.5l-3.5-3.5"/><path d="M16.75 15.75h-6a4 4 0 0 1-4-4v-7"/></g>`),
		g.Group(children),
	)
}