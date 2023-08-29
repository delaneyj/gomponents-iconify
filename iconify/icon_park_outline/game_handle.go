package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameHandle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M32 18H16C9.373 18 4 23.373 4 30s5.373 12 12 12h16c6.627 0 12-5.373 12-12s-5.373-12-12-12Z"/><path stroke-linecap="round" d="M16 26v8m-4-4h8m4-14V9.714h8V4"/><path d="M32 34a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}