package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeRightAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q-2.075 0-3.538-1.463T4 12q0-2.075 1.463-3.538T9 7q1.825 0 3.188 1.137T13.9 11h4.275l-.9-.9Q17 9.825 17 9.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.15.15.213.325t.062.375q0 .2-.062.375t-.213.325l-2.6 2.6q-.275.275-.688.275T17.3 15.3q-.3-.3-.3-.713t.3-.712l.875-.875H13.9q-.35 1.725-1.713 2.863T9 17Z"/>`),
		g.Group(children),
	)
}