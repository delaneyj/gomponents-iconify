package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitTurnDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M37 22H21a8 8 0 0 0-8 8v14"/><circle cx="37" cy="8.944" r="5" transform="rotate(-90 37 8.944)"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 14v29m5-4l-5 5l-5-5m-14 0l-5 5l-5-5"/></g>`),
		g.Group(children),
	)
}