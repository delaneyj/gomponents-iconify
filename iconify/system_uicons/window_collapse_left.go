package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowCollapseLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 15.5v-10a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M5.5 15.5v-10a2 2 0 0 1 2-2h-2c-1 0-2 .895-2 2v10c0 1.105 1 2 2 2h2a2 2 0 0 1-2-2z"/><path d="m10.5 13.5l-3-3l3-3m5 3h-8"/></g>`),
		g.Group(children),
	)
}