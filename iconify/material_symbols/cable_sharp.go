package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CableSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2H3v-5h2V7q0-1.65 1.175-2.825T9 3q1.65 0 2.825 1.175T13 7v10q0 .825.588 1.413T15 19q.825 0 1.413-.588T17 17v-7h-2V5h1V3h4v2h1v5h-2v7q0 1.65-1.175 2.825T15 21q-1.65 0-2.825-1.175T11 17V7q0-.825-.588-1.413T9 5q-.825 0-1.413.588T7 7v7h2v5H8v2H4Z"/>`),
		g.Group(children),
	)
}