package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><circle cx="24" cy="8" r="5" stroke-linejoin="round"/><path d="M5 28s17-20.25 38 0"/><path stroke-linejoin="round" d="M19 28v-3.79S19 19 24 19s5 5.21 5 5.21V32s0 5-5 5s-5-5-5-5v-4Zm10 4l8 5l-6 7M19 32l-8 5l6 7"/></g>`),
		g.Group(children),
	)
}