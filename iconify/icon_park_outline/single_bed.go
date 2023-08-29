package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SingleBed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 12a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v11H8V12ZM6 35v4m36-4v4"/><path d="M29 18H19a3 3 0 0 0-3 3v2h16v-2a3 3 0 0 0-3-3ZM4 26a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v9H4v-9Z"/></g>`),
		g.Group(children),
	)
}