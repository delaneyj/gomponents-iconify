package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBidirectionalLeftRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.457 6.543a1 1 0 0 1 0 1.414L6.414 11h11.172l-3.043-3.043a1 1 0 0 1 1.414-1.414l4.75 4.75a1 1 0 0 1 0 1.414l-4.75 4.75a1 1 0 0 1-1.414-1.414L17.586 13H6.414l3.043 3.043a1 1 0 1 1-1.414 1.414l-4.75-4.75a1 1 0 0 1 0-1.414l4.75-4.75a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}