package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M8.778 17.012c0-.559.453-1.012 1.012-1.012h23.976c.559 0 1.012.453 1.012 1.012V31c0 7.18-5.82 13-13 13s-13-5.82-13-13V17.012Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.778 23h26v8h-26zm13-19v6m-8-4v2m16-2v2"/><path stroke-linecap="round" d="M34.778 34a7 7 0 1 0 0-14"/></g>`),
		g.Group(children),
	)
}