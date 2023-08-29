package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2Zm-.498 3a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5H9.5a.5.5 0 0 0 0-1H8.002V5.5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}