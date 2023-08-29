package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceNeedle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247 32v87h18V32zm-40.4 105l14.5 29l-68 17h205.8l-68-17l14.5-29zm-73 64l-7 14h258.8l-7-14zm-55.88 32l21 14H413.3l21-14H77.72zm63.58 32l8.8 22h211.8l8.8-22zM124 305v18h50.3l36.6 171h18.4l-36.6-171H240v171h32V323h47.3l-36.6 171h18.4l36.6-171H388v-18H124z"/>`),
		g.Group(children),
	)
}