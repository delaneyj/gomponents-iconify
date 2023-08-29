package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Journey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M24 40c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15c-8.284 0-15 6.716-15 15c0 8.284 6.716 15 15 15Z"/><path d="M20 11c1.805 1.008 3.5 2.5 3.358 4.445c-.114 1.555-1.443 1.902-1.721 3.026c-.278 1.124 1.33 2.35-1.39 4.165C18.431 23.846 12.97 26.145 9 24"/><path stroke-linejoin="round" d="M9.5 24.5C6.5 26.388 2.068 31.521 4 35c2.5 4.5 12 .69 23-8S42.23 5.606 42.23 5.606L37 7"/><path d="M26 40s.5-5 4-8s8-3 8-3"/></g>`),
		g.Group(children),
	)
}