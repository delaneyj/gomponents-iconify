package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 18a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/><path d="M11 18h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H8.5A5.5 5.5 0 0 0 3 12.5V16a2 2 0 0 0 2 2h2M8 7l7-3l1 3"/><path d="M13 11.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5zm7 4.5h2"/></g>`),
		g.Group(children),
	)
}