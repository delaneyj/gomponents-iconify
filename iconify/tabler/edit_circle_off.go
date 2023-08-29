package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.507 10.498L9 12v3h3l1.493-1.498m2-2.01l4.89-4.907a2.1 2.1 0 0 0-2.97-2.97L12.5 8.511M16 5l3 3"/><path d="M7.476 7.471A7 7 0 0 0 10 21a7 7 0 0 0 6.53-4.474M3 3l18 18"/></g>`),
		g.Group(children),
	)
}