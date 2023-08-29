package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.105 3a2.25 2.25 0 0 0-1.948 1.125l-3.896 6.75a2.25 2.25 0 0 0 0 2.25l3.896 6.75A2.25 2.25 0 0 0 8.105 21h7.79a2.25 2.25 0 0 0 1.95-1.125l3.895-6.75a2.25 2.25 0 0 0 0-2.25l-3.896-6.75A2.25 2.25 0 0 0 15.895 3h-7.79Z"/>`),
		g.Group(children),
	)
}