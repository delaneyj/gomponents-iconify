package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.085 4H10a2 2 0 0 1 2 2v4.915A1.5 1.5 0 0 0 13 9.5V6a3 3 0 0 0-3-3H6.5a1.5 1.5 0 0 0-1.415 1ZM4.5 5A1.5 1.5 0 0 0 3 6.5v5A1.5 1.5 0 0 0 4.5 13h5a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 9.5 5h-5Z"/>`),
		g.Group(children),
	)
}