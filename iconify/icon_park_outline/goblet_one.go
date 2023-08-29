package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GobletOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m5.736 24.121l4.95 4.95c5.077 5.077 13.308 5.077 18.385 0v0c5.077-5.077 5.077-13.308 0-18.385l-4.95-4.95"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" d="m30 30l6 6"/><ellipse cx="14" cy="14" rx="13" ry="7" transform="rotate(-45 14 14)"/><ellipse cx="38" cy="38" rx="6" ry="3" transform="rotate(-45 38 38)"/></g>`),
		g.Group(children),
	)
}