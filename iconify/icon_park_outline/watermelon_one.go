package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatermelonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 4l17 29.92S36.046 38 24 38S7 33.92 7 33.92L24 4Z"/><circle cx="24" cy="17" r="2" fill="currentColor"/><circle cx="27" cy="23" r="2" fill="currentColor"/><circle cx="21" cy="23" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M41 39.92S36.046 44 24 44S7 39.92 7 39.92"/></g>`),
		g.Group(children),
	)
}