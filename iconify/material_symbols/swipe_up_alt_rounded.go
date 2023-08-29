package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeUpAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-2.075 0-3.538-1.463T7 15q0-1.825 1.137-3.188T11 10.1V5.825l-.9.9Q9.825 7 9.412 7T8.7 6.7q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.213T12 2.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.275.688T15.3 6.7q-.3.3-.713.3t-.712-.3L13 5.825V10.1q1.725.35 2.863 1.713T17 15q0 2.075-1.463 3.538T12 20Z"/>`),
		g.Group(children),
	)
}