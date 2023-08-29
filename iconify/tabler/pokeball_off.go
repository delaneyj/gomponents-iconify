package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PokeballOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.04 16.048A9 9 0 0 0 7.957 3.958m-2.32 1.678a9 9 0 1 0 12.737 12.719"/><path d="M9.884 9.874a3 3 0 1 0 4.24 4.246m.57-3.441a3.012 3.012 0 0 0-1.41-1.39M3 12h6m7 0h5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}