package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 5h16v18H7V5Zm0 0h16V3c0-1.105-.895-2-1.994-2H8.994A1.992 1.992 0 0 0 7 3v2ZM1 8.009C1 6.899 1.898 6 2.998 6H7v12H2.998A2.005 2.005 0 0 1 1 15.991V8.01ZM11 8v10m4-10v10m4-10v10"/>`),
		g.Group(children),
	)
}