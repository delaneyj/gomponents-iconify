package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 12a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z" opacity=".5"/><path stroke-linecap="round" d="M2 12h2m16 0h2M12 4V2m0 20v-2" opacity=".5"/></g>`),
		g.Group(children),
	)
}