package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.984 12.535a9 9 0 1 0-8.431 8.448"/><path d="M12 7v5l3 3m4 1v6m3-3l-3 3l-3-3"/></g>`),
		g.Group(children),
	)
}