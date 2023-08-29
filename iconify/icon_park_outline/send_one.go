package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M42 6L4 20.138l20 3.87L29.005 44L42 6Z"/><path stroke-linecap="round" d="m24.008 24.008l5.657-5.656"/></g>`),
		g.Group(children),
	)
}