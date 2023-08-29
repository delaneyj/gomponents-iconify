package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GpsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12.5 14.959a3 3 0 1 1 2.459-2.459M2 12h2m16 0h2M12 4V2m0 20v-2"/><path d="M8 5.07A8 8 0 1 1 5.07 8"/></g>`),
		g.Group(children),
	)
}