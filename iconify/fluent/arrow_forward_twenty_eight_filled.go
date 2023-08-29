package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.543 6.957a1 1 0 0 1 1.414-1.414l5.75 5.75a1 1 0 0 1 0 1.414l-5.75 5.75a1 1 0 0 1-1.414-1.414L21.586 13H13.75a9 9 0 0 0-9 9a1 1 0 1 1-2 0c0-6.075 4.925-11 11-11h7.836l-4.043-4.043Z"/>`),
		g.Group(children),
	)
}