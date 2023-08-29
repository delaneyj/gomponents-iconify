package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundaboutRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-6.075q-2.2-.35-3.6-2.038T2 9q0-2.5 1.75-4.25T8 3q2.2 0 3.888 1.4T13.924 8h4.25L16.6 6.4L18 5l4 4l-4 4l-1.425-1.4l1.6-1.6H13.9q-.725 0-1.275-.475t-.675-1.2q-.225-1.45-1.35-2.388T8 5Q6.35 5 5.175 6.175T4 9q0 1.475.938 2.6t2.387 1.35q.725.125 1.2.675T9 14.9V21H7Z"/>`),
		g.Group(children),
	)
}