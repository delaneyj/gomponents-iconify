package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvPlugCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10 13.154V21m5-12.615v2.769a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2.77a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2Zm-1.667-2V3M6.667 6.385V3"/><path stroke-linejoin="round" d="M16.667 16L15 19h4l-1.667 3"/></g>`),
		g.Group(children),
	)
}