package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreezingLineColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M40 6H8a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M16.123 6L6 15m20.003-9L6 24M35 6L6 33m13-3L6 42m35-32L29.243 20.852M19 21v21m0-21h23"/></g>`),
		g.Group(children),
	)
}