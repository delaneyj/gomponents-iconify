package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeLeftAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 17q-1.825 0-3.188-1.137T10.1 13H5.825l.9.9q.275.275.275.688t-.3.712q-.275.275-.7.275t-.7-.275l-2.6-2.6q-.15-.15-.213-.325T2.426 12q0-.2.063-.375T2.7 11.3l2.6-2.6q.275-.275.688-.275T6.7 8.7q.3.3.3.713t-.3.712L5.825 11H10.1q.35-1.725 1.713-2.863T15 7q2.075 0 3.538 1.463T20 12q0 2.075-1.463 3.538T15 17Z"/>`),
		g.Group(children),
	)
}