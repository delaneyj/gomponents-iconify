package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 10v12h32l6-6l-6-6H6Z"/><path stroke-linecap="round" d="M23 22v22m0-40v6m-5 34h10"/></g>`),
		g.Group(children),
	)
}