package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16Zm-.5 3a.5.5 0 0 0-.5.5v5l.008.09A.5.5 0 0 0 9.5 11h3l.09-.008A.5.5 0 0 0 12.5 10H10V5.5l-.008-.09A.5.5 0 0 0 9.5 5Z"/>`),
		g.Group(children),
	)
}