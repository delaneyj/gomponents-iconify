package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Computing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="5" y="5" rx="2"/><path d="M8 5V2m8 3V3l1-1m-1 20v-3m-7 3v-3M5 8H2m20 0h-3m3 8h-3M5 16H3l-1 1"/></g>`),
		g.Group(children),
	)
}