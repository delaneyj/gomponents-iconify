package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSignBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10 8v8h28l4-4l-4-4H10Zm28 15v8H10l-4-4l4-4h28Z"/><path stroke-linecap="round" d="M24 31v13m0-28v7m0-19v4m-5 36h10"/></g>`),
		g.Group(children),
	)
}