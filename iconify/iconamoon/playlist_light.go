package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6L3 7.732V4.268L6 6Z"/><path stroke-linecap="round" d="M3 12h18M10 6h11M3 18h18"/></g>`),
		g.Group(children),
	)
}