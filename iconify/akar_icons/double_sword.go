package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleSword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 19.2L3.8 21m.9-7.2l.9 3.6m0 0l3.6.9m-3.6-.9l-2.7 2.7M16.4 3.9l-9 9l.45 2.25l2.25.45l9-9L20 3l-3.6.9Z"/><path d="M22 19.2L20.2 21m-.9-7.2l-.9 3.6m0 0l2.7 2.7m-2.7-2.7l-1.8.45l-1.8.45M9.3 11L4.9 6.6L4 3l3.6.9L12 8.3m.1 5.5l1.8 1.8l2.25-.45l.45-2.25l-1.8-1.8"/></g>`),
		g.Group(children),
	)
}