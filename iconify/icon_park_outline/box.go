package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="36" height="30" x="6" y="12" rx="2"/><path stroke-linecap="round" d="M17.95 24.008h12M6 13l7-8h22l7 8"/></g>`),
		g.Group(children),
	)
}