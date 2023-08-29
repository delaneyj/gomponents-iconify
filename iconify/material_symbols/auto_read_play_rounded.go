package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoReadPlayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.175 13.225l3.9-2.6q.35-.225.35-.625t-.35-.625l-3.9-2.6q-.375-.25-.775-.038T10 7.4v5.2q0 .45.4.663t.775-.038ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Z"/>`),
		g.Group(children),
	)
}