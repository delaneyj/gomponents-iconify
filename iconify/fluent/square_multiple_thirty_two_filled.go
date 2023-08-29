package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMultipleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.113 5.5h2.097A2.5 2.5 0 0 1 12.502 4h11a4.5 4.5 0 0 1 4.5 4.5v11a2.5 2.5 0 0 1-1.502 2.293v2.096a4.502 4.502 0 0 0 3.502-4.389v-11a6.5 6.5 0 0 0-6.5-6.5h-11a4.502 4.502 0 0 0-4.389 3.5ZM6.5 7A4.5 4.5 0 0 0 2 11.5v14A4.5 4.5 0 0 0 6.5 30h14a4.5 4.5 0 0 0 4.5-4.5v-14A4.5 4.5 0 0 0 20.5 7h-14Z"/>`),
		g.Group(children),
	)
}