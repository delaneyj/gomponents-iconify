package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Treadmill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0M3 14l4 1l.5-.5M12 18v-3l-3-2.923L9.75 7"/><path d="M6 10V8l4-1l2.5 2.5l2.5.5m6 12a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1m15-1l1-11l2-1"/></g>`),
		g.Group(children),
	)
}