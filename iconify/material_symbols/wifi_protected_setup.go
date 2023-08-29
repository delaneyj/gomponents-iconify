package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiProtectedSetup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.675 18.75q.425-.825.65-1.738t.225-1.887q0-2-.875-3.663T13.3 8.7L11 11V3h8l-2.3 2.3q1.3 1.175 2.075 2.8t.775 3.525q0 2.275-1.063 4.125t-2.812 3ZM5 21l2.3-2.3q-1.325-1.175-2.087-2.8t-.763-3.525q0-2.275 1.063-4.125t2.837-3q-.425.825-.663 1.738T7.45 8.874q0 2 .888 3.663T10.7 15.3L13 13v8H5Z"/>`),
		g.Group(children),
	)
}