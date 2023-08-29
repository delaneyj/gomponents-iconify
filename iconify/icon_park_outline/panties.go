package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panties(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M44 15c-3-3-5-9-5-9s-5 3.5-15 3.5S9 6 9 6s-1 5-5 9c0 12 17 27 17 27h6s17-15 17-27Z"/><path d="M44 15s-9 1-13 8s-4 19-4 19M4 15s9 1 13 8s4 19 4 19"/></g>`),
		g.Group(children),
	)
}