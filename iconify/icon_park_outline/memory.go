package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M8 6v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V13.61a2 2 0 0 0-.605-1.433l-7.813-7.61A2 2 0 0 0 30.187 4H10a2 2 0 0 0-2 2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M40 21H8m32 8H30m10 7H30m0 8V21M18 44V21m0 12H8"/></g>`),
		g.Group(children),
	)
}