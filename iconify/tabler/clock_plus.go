package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.984 12.535a9 9 0 1 0-8.468 8.45M16 19h6m-3-3v6"/><path d="M12 7v5l3 3"/></g>`),
		g.Group(children),
	)
}