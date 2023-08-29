package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5h9l3 5l-3.308 3.697m-1.883 2.104L12.5 19.5a.7.7 0 0 1-1 0L3 10l2.62-4.368"/><path d="M10 12L8 9.8l.6-1M3 3l18 18"/></g>`),
		g.Group(children),
	)
}