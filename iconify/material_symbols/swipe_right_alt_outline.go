package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeRightAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q-2.075 0-3.538-1.463T4 12q0-2.075 1.463-3.538T9 7q1.825 0 3.188 1.137T13.9 11h4.275L16.6 9.4L18 8l4 4l-4 4l-1.425-1.4l1.6-1.6H13.9q-.35 1.725-1.713 2.863T9 17Zm0-2q1.25 0 2.125-.875T12 12q0-1.25-.875-2.125T9 9q-1.25 0-2.125.875T6 12q0 1.25.875 2.125T9 15Zm0-3Z"/>`),
		g.Group(children),
	)
}