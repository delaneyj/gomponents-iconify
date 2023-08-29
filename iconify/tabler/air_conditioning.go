package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirConditioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 16a3 3 0 0 1-3 3m11-3a3 3 0 0 0 3 3m-7-3v4M3 7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M7 13v-3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v3"/></g>`),
		g.Group(children),
	)
}