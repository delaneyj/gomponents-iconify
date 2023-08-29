package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2Zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3Zm-.498 2a.5.5 0 0 1 .491.41l.009.09V8H9.5a.5.5 0 0 1 .09.992L9.5 9H7.502a.5.5 0 0 1-.492-.41l-.008-.09v-3a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}