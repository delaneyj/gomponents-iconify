package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthernetOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linejoin="round" rx="3"/><path stroke-linejoin="round" d="M19 27h10v6H19zm-5-12h20v12H14z"/><path d="M21 19v-4m6 4v-4"/></g>`),
		g.Group(children),
	)
}