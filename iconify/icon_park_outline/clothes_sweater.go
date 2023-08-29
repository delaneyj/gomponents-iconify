package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesSweater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 19v18m-26 0v7h26v-7m-26 0H4V12l5-5l7-3h16l7 3l5 5v25h-7m-26 0V19m13 2v23M9 7l15 14M16 4l8 17m8-17l-8 17M39 7L24 21"/>`),
		g.Group(children),
	)
}