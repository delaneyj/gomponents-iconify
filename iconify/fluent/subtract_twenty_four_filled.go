package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.996 13H20a1 1 0 1 0 0-2H3.996a1 1 0 1 0 0 2Z"/>`),
		g.Group(children),
	)
}