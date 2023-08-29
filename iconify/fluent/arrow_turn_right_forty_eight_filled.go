package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M30.56 26.56a1.5 1.5 0 0 1-2.12-2.12L34.877 18H17.5a4.5 4.5 0 0 0-4.5 4.5v18a1.5 1.5 0 1 1-3 0v-18a7.5 7.5 0 0 1 7.5-7.5h17.38l-6.44-6.44a1.5 1.5 0 0 1 2.12-2.12l9 9a1.5 1.5 0 0 1 0 2.12l-9 9Z"/>`),
		g.Group(children),
	)
}