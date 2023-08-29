package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M33.778 4h-18v8h18V4Z"/><path stroke-linecap="round" d="M15.366 4h-4.588a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H34.19"/><path stroke-linecap="round" d="m27.777 20l-8 8.001h10.005l-8.004 8"/></g>`),
		g.Group(children),
	)
}