package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M10 38v6h28v-6m0-18v-6L30 4H10v16"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 4v10h10"/><path stroke-linecap="round" d="M16 12h4"/><rect width="40" height="18" x="4" y="20" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 25h6l-6 8h6"/><path stroke-linecap="round" d="M24 25v8m7-8v8"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 25h3.5a2.5 2.5 0 0 1 2.5 2.5v0a2.5 2.5 0 0 1-2.5 2.5H31"/></g>`),
		g.Group(children),
	)
}