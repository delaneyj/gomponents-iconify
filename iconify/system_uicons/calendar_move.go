package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 10.462v-6.03a2 2 0 0 1 2-2h.01l12 .058a2 2 0 0 1 1.99 2v12a2 2 0 0 1-2 2h-.01l-12-.057a2 2 0 0 1-1.99-2V14.5m0-8h16"/><path d="m8.5 15.5l3-3l-3-3m3 3h-10"/></g>`),
		g.Group(children),
	)
}