package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 44c8.284 0 15-6.716 15-15C39 15 24 4 24 4S9 15 9 29c0 8.284 6.716 15 15 15Z" clip-rule="evenodd"/><path d="M9 29c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0 0-9 3-15 0S9 29 9 29Z"/></g>`),
		g.Group(children),
	)
}