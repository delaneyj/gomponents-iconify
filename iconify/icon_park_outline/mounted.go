package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mounted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M9 7h12m6 0h12"/><path stroke-linejoin="round" d="M19 36h-3a3 3 0 0 1-3-3v-9a8 8 0 0 1 8-8h6a8 8 0 0 1 8 8v9a3 3 0 0 1-3 3h-3"/><circle cx="24" cy="7" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 35v1.4a1.6 1.6 0 0 1-1.6 1.6h-6.8a1.6 1.6 0 0 1-1.6-1.6V35a5 5 0 0 1 10 0Z"/><path stroke-linejoin="round" d="M21 38v3a3 3 0 1 0 6 0v-3"/></g>`),
		g.Group(children),
	)
}