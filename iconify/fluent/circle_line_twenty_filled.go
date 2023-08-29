package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.985 10.5H2.015a8 8 0 0 0 15.97 0Zm0-1H2.015a8 8 0 0 1 15.97 0Z"/>`),
		g.Group(children),
	)
}