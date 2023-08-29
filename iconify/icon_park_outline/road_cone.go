package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadCone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M27 6h-6l-1.25 4.5L18.5 15L16 24l-2.5 9l-1.25 4.5L11 42h26l-1.5-5.4L32 24l-2.5-9l-1.25-4.5L27 6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m37 42l-1.5-5.4L32 24l-2.5-9l-1.25-4.5L27 6h-6l-1.25 4.5L18.5 15L16 24l-2.5 9l-1.25 4.5L11 42m26 0H11h26Zm0 0H6h31Zm0 0h5h-5Zm-23.5-9h21M32 24H16m13.5-9h-11m9.75-4.5l7.25 26.1m-23.25.9l7.5-27"/></g>`),
		g.Group(children),
	)
}