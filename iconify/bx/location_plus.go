package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.42 21.815a1.004 1.004 0 0 0 1.16 0C12.884 21.598 20.029 16.44 20 10c0-4.411-3.589-8-8-8S4 5.589 4 9.996c-.029 6.444 7.116 11.602 7.42 11.819zM12 4c3.309 0 6 2.691 6 6.004c.021 4.438-4.388 8.423-6 9.731c-1.611-1.308-6.021-5.293-6-9.735c0-3.309 2.691-6 6-6z"/><path fill="currentColor" d="M11 14h2v-3h3V9h-3V6h-2v3H8v2h3z"/>`),
		g.Group(children),
	)
}