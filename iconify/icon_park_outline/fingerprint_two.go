package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 44V22a6 6 0 0 1 12 0v22"/><path d="M13 44V22c0-6.627 5.373-12 12-12s12 5.373 12 12v22"/><path d="M7 44V22c0-9.941 8.059-18 18-18s18 8.059 18 18v22m-18 0V22"/></g>`),
		g.Group(children),
	)
}