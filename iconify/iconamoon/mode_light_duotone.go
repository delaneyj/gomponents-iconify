package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeLightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="4" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="4" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`),
		g.Group(children),
	)
}