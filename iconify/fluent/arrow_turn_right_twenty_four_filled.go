package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.458 3.293a1 1 0 1 0-1.414 1.414L17.337 7H10.5A4.5 4.5 0 0 0 6 11.5V20a1 1 0 1 0 2 0v-8.5A2.5 2.5 0 0 1 10.5 9h7.335l-2.792 2.791a1 1 0 1 0 1.414 1.415l4.25-4.25a1 1 0 0 0 0-1.413l-4.249-4.25Z"/>`),
		g.Group(children),
	)
}