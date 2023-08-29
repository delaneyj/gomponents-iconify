package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GardenCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 17.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0M6 8v11a1 1 0 0 0 1.806.591L11.5 14.5v.055"/><path d="M6 8h15l-3.5 7l-7.1-.747a4 4 0 0 1-3.296-2.493L4.251 4.63A1 1 0 0 0 3.323 4H2"/></g>`),
		g.Group(children),
	)
}