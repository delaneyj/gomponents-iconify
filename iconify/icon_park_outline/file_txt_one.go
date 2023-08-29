package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTxtOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M10 38v6h28v-6m0-18v-6L30 4H10v16"/><path stroke-linejoin="round" d="M28 4v10h10"/><path d="M16 12h4"/><rect width="40" height="18" x="4" y="20" stroke-linejoin="round" rx="2"/><path stroke-linejoin="round" d="m21 25l6 8m0-8l-6 8m-8-8v8m-3-8h6m19 0v8m-3-8h6"/></g>`),
		g.Group(children),
	)
}