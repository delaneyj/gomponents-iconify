package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" stroke-linejoin="round" rx="2"/><path d="M5 18h38M5 30h38M17 5v38M30 5v38"/></g>`),
		g.Group(children),
	)
}