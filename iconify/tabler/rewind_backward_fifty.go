package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackwardFifty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.007 16.466A6 6 0 0 0 15 6H4m8 9.5v3a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0zM6 20h1.5a1.5 1.5 0 0 0 0-3H6v-3h3"/><path d="M7 9L4 6l3-3"/></g>`),
		g.Group(children),
	)
}