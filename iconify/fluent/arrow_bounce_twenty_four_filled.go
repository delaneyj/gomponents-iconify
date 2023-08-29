package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBounceTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v7a1 1 0 1 0 2 0V9.414l8.043 8.043a1 1 0 0 0 1.414 0l8.25-8.25a1 1 0 0 0-1.414-1.414l-7.543 7.543L5.414 8H10a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}