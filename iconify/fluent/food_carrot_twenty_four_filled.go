package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodCarrotTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.75a.75.75 0 0 0-1.5 0v4.015c-2.026-1.358-4.867-.881-6.293 1.215L2.353 18.786c-.556.818-.45 1.91.255 2.608a2.1 2.1 0 0 0 2.667.23l10.79-7.469A4.454 4.454 0 0 0 17.24 8h4.009a.75.75 0 0 0 0-1.5h-2.69l3.22-3.22a.75.75 0 0 0-1.06-1.06L17.5 5.439V2.75Z"/>`),
		g.Group(children),
	)
}