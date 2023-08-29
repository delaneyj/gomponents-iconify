package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.575q-.2 0-.375-.062T11.3 21.3l-2.6-2.6q-.275-.275-.275-.688T8.7 17.3q.3-.3.713-.3t.712.3l.875.875V13.9q-1.725-.35-2.863-1.713T7 9q0-2.075 1.463-3.538T12 4q2.075 0 3.538 1.463T17 9q0 1.825-1.137 3.188T13 13.9v4.275l.9-.9q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.15.15-.325.213t-.375.062ZM12 12q1.25 0 2.125-.875T15 9q0-1.25-.875-2.125T12 6q-1.25 0-2.125.875T9 9q0 1.25.875 2.125T12 12Zm0-3Z"/>`),
		g.Group(children),
	)
}