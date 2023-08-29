package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleaves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 43s-13 1-20-7S4 4 4 4s24-1 32 5s6 23 6 23"/><path d="M44 44s-11.18-8.449-18-16c-6.82-7.552-10-15-10-15m10 15l1-13m-1 13l-10-1"/></g>`),
		g.Group(children),
	)
}