package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 15.5h-5a.75.75 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0V14h4.25a.75.75 0 0 1 0 1.5ZM14 2C7.372 2 2 7.373 2 14s5.372 12 12 12c6.627 0 12-5.373 12-12S20.627 2 14 2Z"/>`),
		g.Group(children),
	)
}