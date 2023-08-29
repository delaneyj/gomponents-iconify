package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crosshair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 5.25v-3m0 11.5v-3M10.75 8h3M2.25 8h3"/><circle cx="8" cy="8" r="6.25"/></g>`),
		g.Group(children),
	)
}