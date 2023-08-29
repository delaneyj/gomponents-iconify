package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortStaySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 9q-.825 0-1.413-.588T7.5 7q0-.825.588-1.413T9.5 5q.825 0 1.413.588T11.5 7q0 .825-.588 1.413T9.5 9ZM17 22q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm1.675-2.625l.7-.7L17.5 16.8V14h-1v3.2l2.175 2.175ZM3 22V2h13v8.075q-.55.075-1.05.238t-.95.362V4H5v6.525q.45-.275.95-.4T7 10h5q.575 0 1.113.163t1.037.437q-1.275.575-2.212 1.55T10.5 14.4V13h-2v2h-2v2h2v2h1.8q.25.875.7 1.65T12.1 22H3Z"/>`),
		g.Group(children),
	)
}