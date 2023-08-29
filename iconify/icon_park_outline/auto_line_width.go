package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoLineWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M5 40V8"/><path stroke-linejoin="round" d="M24 13L14 35m4-7h12m-6-15l10 22"/><path d="M43 40V8"/></g>`),
		g.Group(children),
	)
}