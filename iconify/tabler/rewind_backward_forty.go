package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackwardForty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.007 16.466A6 6 0 0 0 15 6H4m8 9.5v3a1.5 1.5 0 0 0 3 0v-3a1.5 1.5 0 0 0-3 0zM6 14v2a1 1 0 0 0 1 1h1m1-3v6"/><path d="M7 9L4 6l3-3"/></g>`),
		g.Group(children),
	)
}