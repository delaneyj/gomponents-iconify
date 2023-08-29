package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FEightKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path d="M30 24a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 9a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M21 16h-7v16m0-8h7"/></g>`),
		g.Group(children),
	)
}