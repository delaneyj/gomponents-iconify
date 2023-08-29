package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><circle cx="6" cy="18" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M9 18V5"/><path d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 16V3"/></g>`),
		g.Group(children),
	)
}