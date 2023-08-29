package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageviewSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.325 17.725l1.4-1.4L15.3 13.9q.35-.55.525-1.15T16 11.5q0-1.875-1.313-3.187T11.5 7Q9.625 7 8.312 8.313T7 11.5q0 1.875 1.313 3.188T11.5 16q.65 0 1.25-.175t1.15-.525l2.425 2.425ZM11.5 14q-1.05 0-1.775-.725T9 11.5q0-1.05.725-1.775T11.5 9q1.05 0 1.775.725T14 11.5q0 1.05-.725 1.775T11.5 14ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}