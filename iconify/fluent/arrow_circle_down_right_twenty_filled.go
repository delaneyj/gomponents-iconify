package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm2.851 10.856A.499.499 0 0 1 12.5 13h-5a.5.5 0 0 1 0-1h3.793L7.146 7.854a.5.5 0 1 1 .708-.708L12 11.293V7.5a.5.5 0 0 1 1 0v5a.5.5 0 0 1-.144.351l-.005.005Z"/>`),
		g.Group(children),
	)
}