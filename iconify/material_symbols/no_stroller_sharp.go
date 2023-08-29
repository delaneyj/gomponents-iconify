package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoStrollerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20q0-.825.588-1.413T6 18q.825 0 1.413.588T8 20q0 .825-.588 1.413T6 22Zm11-7.825L11.975 9.15L16.2 4.175q.5-.575 1.125-.875T18.65 3q1.4 0 2.375.975T22 6.35V7h-2v-.65q0-.575-.387-.963T18.65 5q-.25 0-.475.1t-.4.3L17 6.275v7.9ZM15 15l2 2H5.275l4.1-4.825L1.4 4.225L2.8 2.8l18.4 18.4l-1.425 1.4l-7.6-7.6H15Zm1 7q-.825 0-1.413-.588T14 20q0-.825.588-1.413T16 18q.825 0 1.413.588T18 20q0 .825-.588 1.413T16 22ZM10.9 8.075L6.525 3.7q.825-.35 1.7-.525T10 3q1.125 0 2.213.275T14.3 4.1l-3.4 3.975Z"/>`),
		g.Group(children),
	)
}