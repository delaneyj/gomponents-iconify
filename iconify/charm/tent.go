package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.25 14.25L8 10l2.75 4.25"/><path d="m9.75 1.75l-8 12.5h12.5l-8-12.5"/></g>`),
		g.Group(children),
	)
}