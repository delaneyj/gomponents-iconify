package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.98 8.5a6 6 0 0 1-11.96 0h11.96Zm0-1H2.02a6 6 0 0 1 11.96 0Z"/>`),
		g.Group(children),
	)
}