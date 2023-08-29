package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundaboutRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20v-5.075q-2.2-.35-3.6-2.038T2 9q0-2.5 1.75-4.25T8 3q2.2 0 3.888 1.4T13.924 8h4.25l-.9-.9Q17 6.825 17 6.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.15.15.213.325t.062.375q0 .2-.062.375T21.3 9.7l-2.6 2.6q-.275.275-.688.275T17.3 12.3q-.3-.3-.3-.713t.3-.712l.875-.875H13.9q-.725 0-1.275-.475t-.675-1.2q-.225-1.45-1.35-2.388T8 5Q6.35 5 5.175 6.175T4 9q0 1.475.938 2.6t2.387 1.35q.725.125 1.2.675T9 14.9V20q0 .425-.288.713T8 21q-.425 0-.713-.288T7 20Z"/>`),
		g.Group(children),
	)
}