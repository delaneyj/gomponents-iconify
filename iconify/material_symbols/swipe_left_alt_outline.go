package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeLeftAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17q-1.825 0-3.188-1.137T10.1 13H5.825L7.4 14.6L6 16l-4-4l4-4l1.425 1.4l-1.6 1.6H10.1q.35-1.725 1.713-2.863T15 7q2.075 0 3.538 1.463T20 12q0 2.075-1.463 3.538T15 17Zm0-2q1.25 0 2.125-.875T18 12q0-1.25-.875-2.125T15 9q-1.25 0-2.125.875T12 12q0 1.25.875 2.125T15 15Zm0-3Z"/>`),
		g.Group(children),
	)
}