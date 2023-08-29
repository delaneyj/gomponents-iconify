package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m12.433 20.017l8.037 4.68a1.958 1.958 0 0 0 1.051.301a2.241 2.241 0 0 0 .489-.055l7.253-1.977A1 1 0 0 0 30 22V7.996a1 1 0 0 0-1.26-.966l-7.237 1.953l-8.037-4.68a1.992 1.992 0 0 0-1.865-.135L4 7.473V2H2v26a2 2 0 0 0 2 2h26v-2H4v-4.342ZM28 21.234l-6.497 1.75l-8.037-4.68a1.993 1.993 0 0 0-1.865-.135L4 21.475v-4.818l8.433-3.641l8.037 4.68a1.962 1.962 0 0 0 1.05.302a2.211 2.211 0 0 0 .49-.056L28 16.31ZM12.433 6.016l8.037 4.68a2.01 2.01 0 0 0 1.54.245L28 9.308v4.926l-6.497 1.75l-8.037-4.68a1.992 1.992 0 0 0-1.865-.136L4 14.474V9.656Z"/>`),
		g.Group(children),
	)
}