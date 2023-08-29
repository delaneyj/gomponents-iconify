package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.983 12.548a9 9 0 1 0-8.45 8.436M19 22v-6m3 3l-3-3l-3 3"/><path d="M12 7v5l2.5 2.5"/></g>`),
		g.Group(children),
	)
}