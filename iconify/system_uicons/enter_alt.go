package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 13.535l-3-3.035l3-3m7 3h-10"/><path d="M16.5 8.5V5.54a2 2 0 0 0-1.992-2l-8-.032A2 2 0 0 0 4.5 5.5v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-3"/></g>`),
		g.Group(children),
	)
}