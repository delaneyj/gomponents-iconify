package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadcastOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="4" d="M24 15a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-width="4" d="M24 15v8m6 0v10M18 23v10"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 33V23l34 .013V33"/><path stroke-linecap="round" stroke-width="5" d="M41 41v1M7 41v1m11-1v1m12-1v1"/></g>`),
		g.Group(children),
	)
}